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
		a := game.MinValue(genesis, 1)

		a.Print()
		genesis = a
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
