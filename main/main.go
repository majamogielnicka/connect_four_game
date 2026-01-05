package main

import (
	"fmt"
	"connect_four/modes"
)

func main() {
	fmt.Println("1 - PvP (console)")
	fmt.Println("2 - AI vs Player (console)")
	fmt.Println("3 - PvP (GUI)")
	fmt.Println("4 - AI vs Player (GUI)")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		modes.PvP()
	case 2:
		modes.AI_vs_P()
	case 3:
		modes.PvP_GUI()
	case 4:
		modes.AI_vs_P_GUI()
	}
}
