package main

import (
	"fmt"
	"land-occupy/game"
)

func gg() {

	genesis := game.NewGame().CurrentState
	var x, y uint8

	genesis.Print()
	for {
		fmt.Print("x: ")
		_, err1 := fmt.Scanf("%d", &x)
		if err1 != nil {
			fmt.Println("Error reading the first uint8:", err1)
			return
		}

		fmt.Print("y: ")
		_, err2 := fmt.Scanf("%d", &y)
		if err2 != nil {
			fmt.Println("Error reading the second uint8:", err2)
			return
		}
		genesis.Move(x, y)
		genesis.Print()

		acs, _ := genesis.ExpandOpponentActions()
		max := game.State{Counts: [3]int8{-127, -127, -127}}
		for idx := range acs {
			min := game.MinValue(acs[idx], -127, 127)
			if min > max.Utility() {
				max = acs[idx]
			}
		}
		genesis = max
		genesis.Print()
	}

}

func main() {
	gg()
	//genesis := game.NewGame().CurrentState
	//ff := genesis.ExpandOpponentActions()
	//for i := range ff {
	//	sd := ff[i].ExpandOpponentActions()
	//	for j := range sd {
	//		sd[j].Print()
	//	}
	//	break
	//}

}
