# DialARoman Service

## Lab

> Write a Web Service to convert roman to arabic and vice versa.
> The service output should be a JSON document with the following fields:
> o status : http response code
> o number: the arabic or roman number
> o url: the url to run the inverse converter
ie if roman was requested formulate the arabic conversion url or vice versa.


## Usage

> This service is listening on port 9000

1. To convert a number from arabic to a roman literal

    ```shell
    http :9000/roman n==10
    ```

1. To convert a roman literal to arabic

    ```shell
    http :9000/arabic r=="XXV"
    ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)