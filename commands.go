package cli

import (
	"fmt"
	"os/exec"
)

var version = "0.1.1"

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
	cmd := exec.Command("make", "build")
	err := cmd.Run()
	if err != nil {
		Error("Build failed", err)
	} else {
		Success("Done")
	}
}
