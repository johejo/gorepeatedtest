# gorepeatedtest

[![CI](https://github.com/johejo/gorepeatedtest/workflows/ci/badge.svg)](https://github.com/johejo/gorepeatedtest/actions?query=workflow%3Aci)
[![Go Reference](https://pkg.go.dev/badge/github.com/johejo/gorepeatedtest.svg)](https://pkg.go.dev/github.com/johejo/gorepeatedtest)
[![Go Report Card](https://goreportcard.com/badge/github.com/johejo/gorepeatedtest)](https://goreportcard.com/report/github.com/johejo/gorepeatedtest)

Run go test repeatedly to find flaky tests.

By default, race and `GOMAXPROCS` are changed randomly.

## Example

```
gorepeatedtest -d 60s -- testing
```

## License

MIT

## Author

Mitsuo Heijo
