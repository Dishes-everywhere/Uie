package main

import "fmt"

func main() {
	var in string
	var a string
	var yes string
	var in2 string
	var in3 string
	var in4 string
	var c string
	var b string
	b = "uie ip"
	var d string
	d = "exit"
	yes = "yes"
	var no string
	no = "no"
	a = "welcome"
	fmt.Println("Please enter the command(please enter lowercase):")
	fmt.Scan(&in)
	if in == a {
		fmt.Println("Welcome please continue typing:")
		fmt.Scan(&in2)
		if in2 == b {
			fmt.Println("Enter ip:")
			fmt.Scan(&c)
			fmt.Println("Ip is ", c, "?", "\n", "enter yes/no continue:", "\n")
			fmt.Scan(&in3)
			if in3 == yes {
				fmt.Println("Ok,uie has hacked into this ip,you've got control!Please enter lowercase:")
			} else if in3 == no {
				fmt.Println("Enter ip:")
				fmt.Scan(&c)
				fmt.Println("Ip is ", c, "?", "\n", "enter yes/no continue:", "\n")
				fmt.Scan(&in4)
				if in4 == yes {
					fmt.Println("Ok,uie has hacked into this ip,you've got control!")
				} else if in4 == no {
					fmt.Println("Please reboot!!!")
					return
				}
			}
		}
	}
	if in == d {
		fmt.Println("exit")
		return
	}
}
