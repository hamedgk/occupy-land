package main

import (
	"fmt"
	"land-occupy/game"
	"math"
)

func main() {

	genesis := game.NewGame().CurrentState
	var x, y int

	genesis.Print()
	for {
		fmt.Print("x: ")
		_, err1 := fmt.Scanf("%d", &x)
		if err1 != nil {
			fmt.Println("Error reading the first int:", err1)
			return
		}

		fmt.Print("y: ")
		_, err2 := fmt.Scanf("%d", &y)
		if err2 != nil {
			fmt.Println("Error reading the second int:", err2)
			return
		}
		genesis.Move(x, y)
		genesis.Print()

		acs, _ := genesis.Expand()
		maxUtil := math.MinInt
		maxVal := game.State{}
		var min int
		for idx := range acs {
			min = game.MinValue(acs[idx], math.MinInt, math.MaxInt, 9)
			if min > maxUtil {
				//maxUtil = game.Utility(&acs[idx])
				maxUtil = min
				maxVal = acs[idx]
			}
		}
		genesis = maxVal
		genesis.Print()
	}

}
