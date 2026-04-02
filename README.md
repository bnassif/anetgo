# AnetGo
[![MIT licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://raw.githubusercontent.com/bnassif/anetgo/main/LICENSE)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/bnassif/anetgo/release.yml)
![GitHub Release](https://img.shields.io/github/v/release/bnassif/anetgo)
![GitHub Release Date](https://img.shields.io/github/release-date/bnassif/anetgo)

AnetGo is a compiled API wrapper for the Atlantic.net API.

> [!NOTE]
> AnetGo is not affiliated with Atlantic.Net itself. This wrapper has been developed using the public 
[API Documentation](https://www.atlantic.net/docs/api/).

## Overview

This package provides a Unix binary shipped with a .deb package which allows full interaction with the Atlantic.net Cloud API.

Subcommands are used to cleanly sort each API call to its own section

## Installation

Refer to the [installation docs](INSTALLATION.md) on how to install `anetctl` for your respective package manager.

## Usage

To use `anetctl`, you must first have a Cloud account with Atlantic.net, and a valid API key and secret.  
You can signup for an account [here](https://cloud.atlantic.net/signup) if you don't already have one.

To obtain your API keys, navigate [here](https://cloud.atlantic.net/?page=account) once signed in, and scroll to the `API Information` section.

Refer to the [documentation](./docs/anetctl.md) in this repository for more details.

## License
MIT - Feel free to use, extend, and contribute.