# LetsGo! Test Greeter Lab

## Mission

> Using Test first principles, implement a set of tests for your very own greeter function
> that we've written in the fmt lab.

* Use ad-hoc values for name, age, height
* Try out zero-values for name, age, height
* Make use of the assert functions from testify to validate your tests.
* See installation/run steps below
* Practice breaking one of the test on purpose and fixing it.
* Issue a test command to only run one of the tests from the command line!
* Make sure you achieve 100% test coverage using the standard coverage tool.

## Installation

```shell
# Install Assertion and Reflex packages
go get -f -u github.com/cespare/reflex github.com/stretchr/testify
```

## Up And Running...

1. Run the tests

```shell
make test
```

1. Run test coverage

```shell
make cov
```

1. Run auto tests

```shell
make watch
```

---
<img src="assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)