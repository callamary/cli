package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var version = "1.0.0"

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
	cwd, _ := os.Getwd()

	fmt.Println("Use default location? (y/n): ")
	var response string
	fmt.Scanln(&response)

	var loc string
	if response == "y" {
		loc = filepath.Join(cwd, "/app/cli/main.go")
	} else {
		var custom string
		fmt.Println("Enter commands main.go file path: ")
		fmt.Scanln(&custom)
		loc = filepath.Join(cwd, custom)
	}

	outputPath := filepath.Join(cwd, "call")
	cmd := exec.Command("go", "build", "-o", outputPath, loc)
	err := cmd.Run()
	if err != nil {
		Error("Build failed", err)
	} else {
		Success("Done\n")
	}
}
