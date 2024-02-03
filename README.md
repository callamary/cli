# Simple CLI Package for Go Projects
This CLI package is designed to facilitate the integration of command-line interface (CLI) capabilities into Go projects, streamlining the process of implementing CLI commands.

## How to use
Assuming that your CLI commands are located in a commands directory, you can utilize the package as follows:
```
// main.go
func main() {
  // Params: name, handler, short description
  cli.RegisterCommand("hello", hello, "Sample cli command")
  ...

  cli.Run() // put this at the bottom
}

func hello() {
  fmt.Println("Hello World")
}
```

## Console Message
Showing message in terminal will be ease. For example you need to print error message in terminal, you can call:

### Error
```cli.Error("Error", error)```

### Warning
```cli.Warning("Warning")```

### Success
```cli.Success("Success")```


## Build the commands
For the initial installation, you must manually build your executable by using the following command:
``go build -o ./call ./commands/main.go``
After the executable has been generated, you can simply run ``./call build`` to incorporate any newly added commands in subsequent updates.



