<img src="../assets/gophernand.png" align="right" width="100" height="auto"/>

<br/>
<br/>
<br/>

# LetsGo! Dictionary Service Lab

---
## <img src="../assets/lab.png" width="auto" height="32"/> Mission

> And that's a Wrap!! Decorate your word dictionary as a Web Service...

* Make sure to leverage your dictionary/jsondic packages for this lab!
* Ensure your web server is listening on port 4500
* Create a HTTP endpoint **/new_word** to serve out new words from a given dictionary
  * The request should take a query parameter ie dic=words or dic=artists
  * You should issue a JSON response!
  * Add a logger to log out your endpoint calls
* In your implementation load the dictionary from the requested dictionary file
* From your *dictionary* package api, load and pick out a random word
* From your *jsondic* package, use the encoder to encode the JSON response
* Make sure the proper JSON headers are set
* Write the necessary tests to shake out your endpoints!
* Run your application and test it using the dictionaries provided in the assets directory
* BONUS! Hit a classmate dictionary service using watch or a shell loop to make sure
  random words are indeed showing up for each requests ;)

## Expected Output

```text
HTTP/1.1 200 OK
Content-Length: 68
Content-Type: application/json; charset=utf-8
Date: Sat, 10 Feb 2018 19:11:38 GMT

{
    "dictionary": "trump",
    "location": "assets",
    "random_word": "fakenews"
}
```

## OSX Brew Installs (Totally Optional!!)

```shell
# Install watch via homebrew
brew install watch
# Install httpie
brew install httpie
```

## Commands

```shell
# Start your web server
go run main.go
# Hit the new_word url classic
curl http://localhost:4500/new_word?dic=words
# Or... Hit the trump dictionary
curl http://localhost:4500/new_word?dic=trump
# Optionally use watch!
watch curl http://localhost:4500/new_word?dic=words
# Or httpie...
http :4500/new_word dic==trump
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)