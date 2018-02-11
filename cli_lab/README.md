<img src="../assets/gophernand.png" align="right" width="100" height="auto"/>

<br/>
<br/>
<br/>

# LetsGo! CLI with Flags Lab

---
## <img src="../assets/lab.png" width="auto" height="32"/> Mission


> Leveraging the code from the greeter lab, create a program that takes in parameters
> for the user name and age and display the message.

* Create a new directory in your Go Worskspace called *letsgo_greeter*
* Create a file *main.go* in that directory ie $GOPATH/src/github.com/YOUR_GITHUB_HANDLE/letsgo_greeter
* Default user: *No one*
* Default age: *42* of course!
* [Extra Bonus] Print out the program help. Can you explain what's going on with Usage of? and how to fix it?

## Usage

```shell
$ go run main.go               # => Hello, my name is No one. I am 42 years old!"
$ go run main.go -u Fernand    # => Hello, my name is Fernand. I am 42 years old!
$ go run main.go -a 21         # => Hello, my name is No One. I am 21 years old!
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)