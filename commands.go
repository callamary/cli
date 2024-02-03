package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func version() string {
	file, err := os.Open("changelog.md")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	versionRegex := regexp.MustCompile(`## \[([0-9]+\.[0-9]+\.[0-9]+)\]`)

	for scanner.Scan() {
		line := scanner.Text()
		if matches := versionRegex.FindStringSubmatch(line); matches != nil {
			return matches[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	return "version not found"
}

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
	fmt.Printf("%s%s%s", "Callamary", "\033[32m version \033[37m", version()+".\r\nDeveloped by \033[32mElva Hadi Prajatama \033[36m@codacrafts\033[37m\n\n")
}

func build() {
	cmd := exec.Command("make", "build")
	err := cmd.Run()
	if err != nil {
		Error("Build failed", err)
	} else {
		Success("Done")
	}
}
