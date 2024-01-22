# GoSaviynt

[![Build Status][build-status-svg]][build-status-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

`gosaviynt` is a Go language client SDK for the Saviynt Enterprise Identity Cloud (EIC) REST API.

This client SDK is auto-generated using the [OpenAPI Generator codegen project](https://github.com/OpenAPITools/openapi-generator).

The REST API is specificed by the OpenAPI Specification and maintained in the [`github.com/saviynt/saviynt-api-reference`](https://github.com/saviynt/saviynt-api-reference) repository.

Questions or requests for this SDK should be posted via [GitHub issues for this repo](https://github.com/saviynt/gosaviynt/issues).

## Usage

The base page, `gosaviynt`, includes basic methods to retrieve an `*oauth2.Token` and `*http.Client`. These can be used with the auto-generated client SDK in the `saviynt` package.

* An example of using `gosaviynt.FetchToken()` is availabe in [`cmd/client`](cmd/client/main.go).
* An example of using the generated client is available in [`cmd/runtimecontrolsdata`](cmd/runtimecontrolsdata/main.go)

 [used-by-svg]: https://sourcegraph.com/github.com/saviynt/gosaviynt/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/saviynt/gosaviynt?badge
 [build-status-svg]: https://github.com/saviynt/gosaviynt/workflows/test/badge.svg
 [build-status-url]: https://github.com/saviynt/gosaviynt/actions/workflows/test.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/saviynt/gosaviynt
 [goreport-url]: https://goreportcard.com/report/github.com/saviynt/gosaviynt
 [codeclimate-status-svg]: https://codeclimate.com/github/saviynt/gosaviynt/badges/gpa.svg
 [codeclimate-status-url]: https://codeclimate.com/github/saviynt/gosaviynt
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/saviynt/gosaviynt
 [docs-godoc-url]: https://pkg.go.dev/github.com/saviynt/gosaviynt
 [loc-svg]: https://tokei.rs/b1/github/saviynt/gosaviynt
 [repo-url]: https://github.com/saviynt/gosaviynt
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/saviynt/gosaviynt/blob/master/LICENSE
