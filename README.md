<h1 align="center">Callamary CLI</h3>
  <p align="center">
    A simple CLI package for Golang project
    <br/>
    <br/>
    <br/>
    .
    <a href="https://github.com/callamary/cli/issues">Report Bug</a>
    .
    <a href="https://github.com/callamary/cli/issues">Request Feature</a>
  </p>
</p>

<p align="center">
  <a href="https://github.com/callamary/cli/actions"><img src="https://github.com/callamary/cli/actions/workflows/go.yml/badge.svg" alt="Build Status"></a>
  <img src="https://img.shields.io/github/downloads/callamary/cli/total?color=dark-blue" alt="Downloads">
  <img src="https://img.shields.io/github/contributors/callamary/cli?color=dark-green" alt="Contributors">
  <img src="https://img.shields.io/github/issues/callamary/cli" alt="Issues">
  <img src="https://img.shields.io/badge/license-MIT-green" alt="License">
</p>


<br />

## Simple CLI Package for Go Projects
This CLI package is designed to facilitate the integration of command-line interface (CLI) capabilities into Go projects, streamlining the process of implementing CLI commands.

### How to use
Assuming that your CLI commands are located in a commands directory, you can utilize the package as follows:
```go
/* commands/main.go */

func main() {

  cli.Command("hello", hello, "Greeting CLI command")
  ...

  cli.Run() // call run function at the bottom
}

func hello() {
  fmt.Println("Hello World")
}
```

### Console Message
Showing message in terminal will be ease. For example you need to print error message in terminal, you can call:

#### Error
```go 
  cli.Error("Error", error)
```

#### Warning
```go 
  cli.Warning("Warning")
```

#### Success
```go 
  cli.Success("Success")
```


### Build the commands
For the initial installation, you must manually build your executable by using the following command:
```bash 
build -o ./call ./commands/main.go
```
After the executable has been generated, you can simply run 
```bash 
./call build
```
to incorporate any newly added commands in subsequent updates.


### License
Callamary CLI is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).



