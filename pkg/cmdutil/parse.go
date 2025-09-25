package cmdutil

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"sort"

	"github.com/bnassif/anetgo/pkg/api"
)

var itemKeyPattern = regexp.MustCompile(`^(\d+item|item)$`)

func HandleRequest(client *api.Client, action string, params map[string]string, raw bool, rootKey string, nestedKey string) {
	var resp string
	var err error

	if raw {
		// Send the request
		resp, err = client.Request(action, params)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(resp)
	} else {
		var data []byte
		data, err = client.RequestRaw(action, params)
		if err != nil {
			log.Fatal(err)
		}

		// Send the request
		resp, err = NormalizeResponse(data, rootKey, nestedKey)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(resp)
	}
}

func NormalizeResponse(data []byte, rootKey, nestedKey string) (string, error) {
	var root map[string]any
	if err := json.Unmarshal(data, &root); err != nil {
		return "", fmt.Errorf("unmarshal root: %w", err)
	}

	// Extract the root object
	obj, ok := root[rootKey]
	if !ok {
		return "", fmt.Errorf("root key %q not found", rootKey)
	}

	objMap, ok := obj.(map[string]any)
	if !ok {
		return "", fmt.Errorf("root value is not an object")
	}

	// If a nested key is given, dive in
	if nestedKey != "" {
		if v, ok := objMap[nestedKey]; ok {
			obj = v
		} else {
			return "", fmt.Errorf("nested key %q not found", nestedKey)
		}
	} else {
		obj = objMap
	}

	// 🔑 Unwrap single-key nested maps (like return → snapshot → {...})
	for {
		m, ok := obj.(map[string]any)
		if !ok || len(m) != 1 {
			break
		}
		for _, v := range m {
			obj = v
			break
		}
	}

	// Normalize: handle different container types
	switch v := obj.(type) {
	case []any: // already a list
		if len(v) == 1 {
			normalized, err := json.Marshal(v[0])
			return string(normalized), err
		}
		normalized, err := json.Marshal(v)
		return string(normalized), err

	case map[string]any: // mapping
		// Detect if keys follow multi-family pattern (like 1item + 1instance)
		grouped := make(map[string]map[string]any)
		multiFamily := true

		for k, val := range v {
			matches := regexp.MustCompile(`^(\d+)([A-Za-z].*)$`).FindStringSubmatch(k)
			if len(matches) == 3 {
				num, subkey := matches[1], matches[2]
				if _, ok := grouped[num]; !ok {
					grouped[num] = make(map[string]any)
				}
				// val must be a map to merge
				if submap, ok := val.(map[string]any); ok {
					for sk, sv := range submap {
						grouped[num][sk] = sv
					}
				} else {
					grouped[num][subkey] = val
				}
			} else {
				multiFamily = false
				break
			}
		}

		if multiFamily && len(grouped) > 0 {
			// Flatten groups into slice
			var keys []string
			for k := range grouped {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			var merged []any
			for _, k := range keys {
				merged = append(merged, grouped[k])
			}

			if len(merged) == 1 {
				normalized, err := json.Marshal(merged[0])
				return string(normalized), err
			}
			normalized, err := json.Marshal(merged)
			return string(normalized), err
		}

		// Existing Xitem detection logic...
		allItemKeys := true
		var keys []string
		for k := range v {
			if !itemKeyPattern.MatchString(k) {
				allItemKeys = false
				break
			}
			keys = append(keys, k)
		}

		if allItemKeys {
			// Sort keys: "item" first, then 1item, 2item...
			sort.Slice(keys, func(i, j int) bool {
				if keys[i] == "item" {
					return true
				}
				if keys[j] == "item" {
					return false
				}
				return keys[i] < keys[j]
			})

			var list []any
			for _, k := range keys {
				list = append(list, v[k])
			}
			if len(list) == 1 {
				normalized, err := json.Marshal(list[0])
				return string(normalized), err
			}
			normalized, err := json.Marshal(list)
			return string(normalized), err
		}

		// Otherwise, return the map as-is
		normalized, err := json.Marshal(v)
		return string(normalized), err

	default: // single object, string, number, etc.
		normalized, err := json.Marshal(v)
		return string(normalized), err
	}
}
