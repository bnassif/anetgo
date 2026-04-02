# Installation
The `anetctl` program is packaged for many Linux distros.  
Please refer to the sections below for your respective package manager.

All examples below assume you've downloaded a packaged release from the [releases page](https://github.com/bnassif/anetgo/releases).

> [!INFO]
> Windows packaging is not provided as of this writing. Windows users are encouraged to build the program manually.

## DEB

### Verification

```bash
# Download the PGP key and import it
wget 'https://github.com/bnassif/anetgo/blob/main/build/keys/pgp-public.key' -o anetctl-pgp.key
gpg --import anetctl-pgp.key

# Verify the signature
dpkg-sig --verify path/to/file.deb
```


### Installation

```bash
# Install the package
dpkg -i path/to/file.deb
```

## RPM

### Verification

```bash
# Install the PGP Public Key
rpm --import 'https://github.com/bnassif/anetgo/blob/main/build/keys/pgp-public.key'

# Verify the signature
rpm --checksig path/to/file.rpm
```

### Installation

```bash
# Install the package
rpm -i path/to/file.rpm
```

## APK

### Installation

```bash
# Install the package
apk add --allow-untrusted path/to/file.apk
```

# Build from Source
Building a package requires:

- [Golang v1.23+](https://go.dev/doc/install)
- [`nFPM`](https://nfpm.goreleaser.com/docs/install/)

```bash
# Clone and enter the project
git clone https://github.com/bnassif/anetctl.git
cd ./anetctl/

# Build just a binary, saved to 'build/bin/anetctl'
make bin

# Make a package for your respective package manager, saved to dist/
make package_deb
make package_rpm
make package_apk
```