# AnetGo
AnetGo is a compiled API wrapper for the Atlantic.net API.

> ***NOTE**: AnetGo is not affiliated with Atlantic.Net itself. This wrapper has been developed using the public API documentation.*

[API Documentation](https://www.atlantic.net/docs/api/)


## Overview

This package provides a Unix binary shipped with a .deb package which allows full interaction with the Atlantic.net Cloud API.

Subcommands are used to cleanly sort each API call to its own section

## Usage

To use `anetctl`, you must first have a Cloud account with Atlantic.net, and a valid API key and secret.  
You can signup for an account [here](https://cloud.atlantic.net/signup) if you don't already have one.

To obtain your API keys, navigate [here](https://cloud.atlantic.net/?page=account) once signed in, and scroll to the `API Information` section.

---

For documentation on usage refer to the [documentation](./docs/anetctl.md) in this repository.

## Build from Source
Building a package requires:

- Golang v1.23+
- `dpkg-deb`

Only Debian dpkg packaging is supported currently.

```bash
git clone https://github.com/bnassif/anetctl.git

# Enter the project's dir
cd ./anetctl/

make release
```

The resulting binary and deb package will be output to the ./artifact directory by default.

To override this, pass the `OUT_ROOT` parameter to make:

```bash
make OUT_ROOT=/tmp/anetctl/ release
```

## License
MIT - Feel free to use, extend, and contribute.