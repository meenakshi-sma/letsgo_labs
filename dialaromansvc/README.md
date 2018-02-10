# LetsGo! DialARoman Service Lab

## Mission

> Going Full Monty! Write a web Service to convert an arabic number to a roman literal.

* The service output should be a JSON document with the following fields:
  * status : http response code
  * result: the roman literal
  * url: the url to run the inverse converter
* Use port 9000 your service
* Use the following roman calculator package:
  [Roman](https://github.com/imhotepio/letsgo_labs/tree/master/roman)
* Make sure all your tests pass!
* Ensure you send an error code back if the arabic number is > 3999
* Make sure code coverage is good > 80%!
* BONUS! Implement the reverse calculator ie from roman -> arabic


## Usage

> This service is listening on port 9000

1. To convert a number from arabic to a roman literal

    ```shell

    curl -XGET http://localhost:3000/roman?n=10
    ```

1. To convert a roman literal to arabic

    ```shell
    curl -XGET http://localhost:3000/arabic?g=X
    ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)