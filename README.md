# consent-receipt-go

[![godoc](https://godoc.org/github.com/adaptant-labs/consent-receipt-go?status.svg)](http://godoc.org/github.com/adaptant-labs/consent-receipt-go)
[![Build Status](https://travis-ci.com/adaptant-labs/consent-receipt-go.svg?branch=master)](https://travis-ci.com/adaptant-labs/consent-receipt-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/adaptant-labs/consent-receipt-go)](https://goreportcard.com/report/github.com/adaptant-labs/consent-receipt-go)

CLI utilities and Go library for generating and validating Consent Receipts in line with the Kantara Initiative
[Consent Receipt Specification](https://kantarainitiative.org/confluence/display/infosharing/Consent+Receipt+Specification).

`consent-receipt-go` implements `KI-CR-v1.1.0` of the consent receipt specification.

## Installing

`consent-receipt-go` makes use of Go modules and was developed and tested with Go `1.12` and `1.13`, respectively.
Installation otherwise follows the same general `go get` pattern:

```
$ go get github.com/adaptant-labs/consent-receipt-go
```

while upgrading can be accomplished with:

```
$ go get -u github.com/adaptant-labs/consent-receipt-go
```

### Setting up Bash Completion

In the case where Bash completion for CLI sub-commands is desired, the CLI app contains a built-in script generator.
To install this system-wide:

```
# consent-receipt-go completion /etc/bash_completion.d/consent-receipt-go
```

and re-source the active user's `$HOME/.bashrc` or simply log out and back in.

By default, when no output file is provided to the completion command, the active completion script is written to
stdout for run-time resolution in a user's `$HOME/.bashrc`.

## Configuration

Configuration is done through a `$HOME/.consent-receipt-go.toml`.

### Configuration for Validation

At present, validation of the JWT token requires knowledge of the signing key. A minimal configuration for this
is:

```toml
[config]

signing-key = "my-totally-secret-key"
```

### Configuration for Generation

Generation of Consent Receipts, on the other hand, requires one (or more) `Data Controller` definition(s). An example
of this is:

```toml
[config]

signing-key = "my-totally-secret-key"
privacy-policy = "https://www.adaptant.io/privacy-policy"

[[controller]]

name = "Adaptant Solutions AG"
contact = "Paul Mundt"
phone = "49-89-904101300"
email = "compliance@adaptant.io"
url = "https://www.adaptant.io"

        [controller.address]

        country = "DE"
        city = "Deisenhofen"
        region = "BY"
        postalcode = "82041"
        address = "Bahnhofstr. 36"

[[controller]]

name = "Additional Data Controller, Inc."
contact = "John Smith"
phone = "1-800-123-4567"
email = "example@example.com"

        [controller.address]

        country = "US"
        city = "San Francisco"
        region = "CA"
        postalcode = "94105"
```

## Usage

```shell script
$ consent-receipt-go --help
consent-receipt-go is a CLI app and library for generating and
validating Consent Receipts pursuant to the Kantara Initiative
Consent Receipt Specificiation (KI-CR-v1.1.0)

Usage:
  consent-receipt-go [command]

Available Commands:
  completion  Generate bash completion script
  generate    Generate a new Consent Receipt
  help        Help about any command
  validate    Validate a Consent Receipt

Flags:
      --config string   config file (default is $HOME/.consent-receipt-go.toml)
  -h, --help            help for consent-receipt-go
      --version         version for consent-receipt-go

Use "consent-receipt-go [command] --help" for more information about a command.
```

Each of the `generate` and `validate` sub-commands further allow specification of whether a Consent Receipt is provided
as a `JWT Token` or as a `JSON file`, as below:

```shell script
Usage:
  consent-receipt-go generate [command]

Available Commands:
  receipt     Generate a JSON-based Consent Receipt
  token       Generate a new JWT token

Usage:
  consent-receipt-go validate [command]

Available Commands:
  receipt     Validate a JSON-based Consent Receipt
  token       Validate a JWT-encoded Consent Receipt token
```


## API Documentation

Online API documentation is provided through `godoc`, and can be accessed directly on the
[package entry](https://godoc.org/github.com/adaptant-labs/consent-receipt-go/api) in the godoc package repository.

## Features and bugs

Please file feature requests and bugs at the [issue tracker][tracker].

[tracker]: https://github.com/adaptant-labs/consent-receipt-go/issues

## Acknowledgements

This project has received funding from the European Union’s Horizon 2020 research and innovation programme under grant
agreement No 731678.

## License

Licensed under the terms of the Apache 2.0 license, the full version of which can be found in the
[LICENSE](https://raw.githubusercontent.com/adaptant-labs/consent-receipt-go/master/LICENSE) file included in the
distribution.
