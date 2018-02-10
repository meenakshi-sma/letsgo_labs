
<img src="../assets/gophernand.png" align="right" width="100" height="auto"/>

<br/>
<br/>
<br/>

# LetsGo! Dictionary CLI Lab

---
## <img src="../assets/lab.png" width="auto" height="32"/> Mission

> Implement a Dictionary CLI to call your new Dictionary HTTP service

* Implement a CLI that takes the following options:
  * -dic: dictionary name
  * -host: a host location where the dictionary is running
  * -port: port where your service is running
* The CLI should make an HTTP call to your running Dictionary Service
* BONUS! Add a hostname flag and call your classmate Dictionary Web Service!

### Directions

```shell
go run main.go # Starts your Dictionary Web Server
```

### Sample Output

```shell
dictionary_cli --dir assets --dic words
# => [DialADic] The word of the day is `director
dictionary_cli --dir assets --dic trump
# => [DialADic] The word of the day is `huge
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)