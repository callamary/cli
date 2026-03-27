package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var version = "2.0.2"

func about() {
	fmt.Println(`
       *. :         @..:      -          -      .
    :        .              .
   .  :+##*=  :   .  +* .#=  .   .         :   .  .   :   .  .  - .  -
  .  *%*---=   ...   #@..@+   ..   ..  . .  ...    ..    . ..      .  .
    -@+       =###*. #%..@+  +###+  +%###*##%+  +###= -%##* *#. =%. .
    :@*.      .--+@= #%..@+  :--*@. #@- *@- +@: :--#@.=@+   :%#-%*  :
  :  -%%*+*#- #%=#@= #@..@* .%#=%@: #@: *@: +@:.%#+%@.=@=    :%@%. .
   :   -=+=:  :==-+. -=  =:  -+=-=  -=  :+. :=. -+=-= .+.  .  *%-
     .       .                     .-    .                  .  #%:  *
                                                            .  -.  *
                                                             .   . `)
	fmt.Printf("%s%s%s", "Callamary", "\033[32m version \033[37m", version+".\r\nDeveloped by \033[32mElva Hadi Prajatama \033[36m@codacrafts\033[37m\n\n")
}

func build() {
	cwd, err := os.Getwd()
	if err != nil {
		Error("Failed to get working directory", err)
		return
	}

	fmt.Println("Use default source location /app/cli/main.go? (y/n): ")
	var response string
	fmt.Scanln(&response)

	var loc string
	if response == "y" {
		loc = "/app/cli/main.go"
	} else {
		var custom string
		fmt.Println("Enter commands source file path: ")
		fmt.Scanln(&custom)
		loc = filepath.Join(cwd, custom)
	}

	outputPath := filepath.Join(cwd, "call")
	cmd := exec.Command("go", "build", "-o", outputPath, loc)
	buildErr := cmd.Run()
	if buildErr != nil {
		Error("Build failed", buildErr)
	} else {
		Success("Done")
	}
}
