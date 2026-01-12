package main

import (
	"connect_four/modes"
	"fmt"
)

func main() {
	fmt.Println("1 - PvP (console)")
	fmt.Println("2 - AI vs Player (console)")
	fmt.Println("3 - PvP (GUI)")
	fmt.Println("4 - AI vs Player (GUI)")
	fmt.Println("5 - AI vs AI for different depths (GUI)")

	var choice int
	var depth int
	var depthO, depthX int

	fmt.Scan(&choice)

	if choice == 2 || choice == 4 {
		var level int
		fmt.Println("Choose difficulty level:")
		fmt.Println("1 for EASY quick yet effective")
		fmt.Println("2 for MEDIUM takes 3 seconds harder to beat")
		fmt.Println("3 for HARD takes 5 seconds, the hardest to beat")
		fmt.Scan(&level)
		depth = 0
		switch level {
		case 1:
			depth = 5
		case 2:
			depth = 7
		case 3:
			depth = 8
		}
	}
	if choice == 5 {
		fmt.Println("Type depth for first (5-8)")
		fmt.Scan(&depthO)
		fmt.Println("Type depth for second (5-8)")
		fmt.Scan(&depthX)
	}

	switch choice {
	case 1:
		modes.PvP()
	case 2:
		modes.AI_vs_P(depth)
	case 3:
		modes.PvP_GUI()
	case 4:
		modes.AI_vs_P_GUI(depth)
	case 5:
		modes.AI_vs_AI_GUI(depthX, depthO)
	}
}
