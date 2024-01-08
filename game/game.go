package game

import (
	"math"
)

type Game struct {
	CurrentState State
}

func NewGame() Game {
	state := State{Turn: Red}
	//var randx, randy = rand.Intn(int(N)), rand.Intn(int(N))
	state.Move(N/2, N/2-1)
	//randx, randy = rand.Intn(int(N)), rand.Intn(int(N))
	state.Move(N/2-1, N/2)
	return Game{CurrentState: state}
}

func MinValue(state State, alpha, beta, depth int) int {
	acs, isterm := state.Expand()
	if isterm || depth == 0{
		return Utility(&state)
	}
	var min int = math.MaxInt
	for idx := range acs {
		term := MaxValue(acs[idx], alpha, beta, depth-1)
		if min > term {
			min = term
		}
		if min <= alpha {
			return min
		}
		if min < beta {
			beta = min
		}
	}
	return min
}

func MaxValue(state State, alpha, beta, depth int) int {
	acs, isterm := state.Expand()
	if isterm || depth == 0{
		return Utility(&state)
	}
	var max int = math.MinInt
	for idx := range acs {
		term := MinValue(acs[idx], alpha, beta, depth-1)
		if max < term {
			max = term
		}
		if max >= beta {
			return max
		}
		if max > alpha {
			alpha = max
		}
	}
	return max
}
