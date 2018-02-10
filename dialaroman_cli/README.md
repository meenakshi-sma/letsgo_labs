# LetsGo! DialARomanCLI Lab

## Mission

> Implement a DialARomanCLI to call DialARoman service using a CLI app

* Run your program by passing either -a for arabic or -g for roman
  * If an arabic number is provided, retrieve the roman numeral via HTTP
  * If a roman numeral is provided, fetch the arabic representation instead
* The CLI must make HTTP calls back to your running DialARoman web service
* ExtraCredit - call your classmate DialARoman Web Service ;-)

### Directions

```shell
go run main.go # Starts your DialARoman web service
```

### Sample Output

```shell
dialaroman_cli -a 10 # => [DialARoman] Roman numeral for 10 is `X
dialaroman_cli -g X  # => [DialARoman] Arabic number for X is `10
```

---
<img src="..//assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)