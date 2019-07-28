# go-hip: HIP in Golang

Package hip provides simple and painless handling of HIP(Host Identity Protocol), implemented in the Go Programming Language.


[![CircleCI](https://circleci.com/gh/wmnsk/go-hip.svg?style=shield)](https://circleci.com/gh/wmnsk/go-hip)[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fwmnsk%2Fgo-hip.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fwmnsk%2Fgo-hip?ref=badge_shield)

[![GolangCI](https://golangci.com/badges/github.com/wmnsk/go-hip.svg)](https://golangci.com/r/github.com/wmnsk/go-hip)
[![GoDoc](https://godoc.org/github.com/wmnsk/go-hip?status.svg)](https://godoc.org/github.com/wmnsk/go-hip)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/wmnsk/go-hip/blob/master/LICENSE)

## Features

* (to be updated)

## Getting Started

### Prerequisites

The following packages should be installed before starting.

```shell-session
go get -u github.com/pkg/errors
go get -u github.com/google/go-cmp/cmp
go get -u github.com/pascaldekloe/goe/verify
```

### Running examples

Not available yet!

## Supported Features

### Packets

Not available yet!

### Parameters

Parameters defined in 5.2. HIP Parameters, RFC 7401 are under its implementation.

| ID    | Parameter Type       | Supported? |
|-------|----------------------|------------|
| 129   | R1Counter            | Yes        |
| 257   | Puzzle               | Yes        |
| 321   | Solution             | Yes        |
| 385   | Seq                  | Yes        |
| 449   | Ack                  | Yes        |
| 511   | DHGroupList          | Yes        |
| 513   | DiffieHellman        | Yes        |
| 579   | HIPCipher            | Yes        |
| 641   | Encrypted            |            |
| 705   | HostID               | Yes        |
| 715   | HITSuiteList         | Yes        |
| 768   | Cert                 |            |
| 832   | Notification         | Yes        |
| 897   | EchoRequestSigned    | Yes        |
| 961   | EchoResponseSigned   | Yes        |
| 2049  | TransportFormatList  | Yes        |
| 61505 | HIPMAC               | Yes        |
| 61569 | HIPMAC2              | Yes        |
| 61633 | HIPSignature2        | Yes        |
| 61697 | HIPSignature         | Yes        |
| 63661 | EchoRequestUnsigned  | Yes        |
| 63425 | EchoResponseUnsigned | Yes        |

## Disclaimer

This is still an experimental project. Any part of implementations(including exported APIs) may be changed before released as v1.0.0.

## Author(s)

Yoshiyuki Kurauchi ([Twitter](https://twitter.com/wmnskdmms) / [LinkedIn](https://www.linkedin.com/in/yoshiyuki-kurauchi/))

_I'm always open to welcome co-authors! Please feel free to talk to me._

## LICENSE

[MIT](https://github.com/wmnsk/go-hip/blob/master/LICENSE)


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fwmnsk%2Fgo-hip.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fwmnsk%2Fgo-hip?ref=badge_large)