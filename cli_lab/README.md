<img src="../assets/gophernand.png" align="right" width="100" height="auto"/>

<br/>
<br/>
<br/>

# LetsGo! CLI Flags Lab

---
## <img src="../assets/lab.png" width="auto" height="32"/> Mission


> Leveraging the code from the greeter lab, create a program that takes in parameters
> for the user name and age and display the message.

* Clone the Git repository from the template link into your Go workspace (Skip if already cloned!)
  * For Example: cd $GOPATH/src/github.com/YOUR_GIT_USERNAME
  * git clone https://github.com/imhotepio/letsgo_labs.git
* cd *cli_lab*
* Edit the file *main.go* and add your code
  * Note: Default user: *No one*
  * Note: Default age: *42* of course!
* Verify your application is running correctly with/without args and CLI help is displaying the correct information
<br/>
* BONUS! Print out the program help. Can you explain what's going on with Usage of? and how to fix it?

### Setup

```shell
# Navigate to your own Go workspace
$ cd $GOPATH/src/gihub.com/MY_GITHUB_HANDLE
# Clone Labs Repo if not done already
$ git clone https://github.com/imhotepio/letsgo_labs.git
# Lab dir
cd cli_lab
# Get cracking!!
```

### Sample Output

```shell
go run main.go               # => Hello, my name is No one. I am 42 years old!"
go run main.go -u Fernand    # => Hello, my name is Fernand. I am 42 years old!
go run main.go -a 21         # => Hello, my name is Nobody. I am 21 years old!
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)