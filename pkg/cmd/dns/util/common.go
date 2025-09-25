package util

import "strconv"

func ZoneIdOrName(arg string) (key string, value string) {
	if _, err := strconv.Atoi(arg); err == nil {
		return "zone_id", arg
	}

	return "domain_name", arg
}
