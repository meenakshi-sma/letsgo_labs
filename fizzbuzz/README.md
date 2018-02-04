# LetsGo! Testing Fizzbuzz Lab...

## Mission

> Write and test a FizzBuzz game!

* In a FizzBuzz Game, children count up from 1
  * If the number is a multiple of 3, they say **Fizz**
  * If the number is a multiple of 5, they say **Buzz**
  * For multiple of both 3 and 5 they say **FizzBuzz**
  * Otherwise they say the number
* Implement a compute function to translate a number to its FizzBuzz counterpart
* Test your compute function using table testing technique
* Ensure you've achieved complete code coverage!


## Installation

```shell
# Install Assertion and Reflex packages
go get -f -u github.com/cespare/reflex github.com/stretchr/testify
```


## Up And Running

### Run your tests

```shell
go test
```

### Check coverages

```shell
go test -coverprofile=cov.out && go tool cover -html=cov.out
```

### Watch your tests

```shell
reflex -g '*.go' go test
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)