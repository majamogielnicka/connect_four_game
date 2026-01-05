package main

import (
	"connect_four/modes"
	"fmt"
)


func main() {
	fmt.Println("-----Choose mode-------\nType 1 for Player vs Player\nType 2 for AI vs Player")
	var mode_choice int
	fmt.Scan(&mode_choice)
	
	if mode_choice==1{
		modes.PvP()
	} else if mode_choice == 2{
		modes.AI_vs_P()
	}
}