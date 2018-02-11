<img src="../assets/gophernand.png" align="right" width="100" height="auto"/>

<br/>
<br/>
<br/>

# LetsGo! Test Greeter Lab

## Mission

> Congratulations! You've just got a new gig working at Google!!
> Your PM says we have to speed up the initial greeting page load that's
> being hit by billions of users! The page greets the users with
> a user name and age.

* Using Benchmark tests implement a second version of the greeter and compare
  with your initial version.
* Make sure your new implementation works the same as the old one
* TIP! Another way to concatenate strings is by using the + operator.
* What are your findings?

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