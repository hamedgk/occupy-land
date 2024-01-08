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
		minUtil := math.MaxInt
		minVal := game.State{}
		var max int
		for idx := range acs {
			max = game.MinMax(acs[idx], math.MinInt, math.MaxInt, 9)
			if max < minUtil {
				//maxUtil = game.Utility(&acs[idx])
				minUtil = max
				minVal = acs[idx]
			}
		}
		genesis = minVal
		genesis.Print()
	}

}
