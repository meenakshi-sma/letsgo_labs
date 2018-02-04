# LetsGo! Test Greeter Lab

## Mission

> Congratulations! You've just got a new gig working at Google!!
> Your PM says we have to speed up the initial greeting page load that's
> being hit by millions of users. The greeting page greets the users with
> a user name and age. Using benchmarking tests implement a second version
> of the greeter and compare with your initial version. You will need to
> ensure that your new function works just like the old one using your
> tests and then run your benchmarks.

* Use the fmt.Sprintf formatter for your initial implementation
* Tip! Another way to concatenate strings is by using the + operator.
    ```go
    s = "Fred" + " is cool!"       // => Fred is cool!
    s += " And to boot handsome!!" // => Fred is cool! And to boot handsome!!
    ```
* What are your findings?
* NOTE! To get more accurate measurements testing.B.ResetTimer() allows to reset your
  benchmark timer so that setup does not impact measurements. Use it!
* NOTE! You will need to leverage a function Itoa to convert an integer to a string.

## Installation

```shell
# Install Assertion and Reflex packages
go get -f -u github.com/cespare/reflex github.com/stretchr/testify
```

## Up And Running...

1. Run the tests

```shell
go test
```

1. Run Benchmarks

```shell
go test --run xxx --bench Greet --benchmem --benchtime=3
```

1. Run auto tests

```shell
make watch
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)