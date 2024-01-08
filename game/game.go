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
		term := MinMax(acs[idx], alpha, beta, depth-1)
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
		term := MinMax(acs[idx], alpha, beta, depth-1)
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

func MinMax(state State, alpha, beta, depth int) int{
	_, isterm := state.Expand()
	if isterm || depth == 0{
		return Utility(&state)
	}
	var val int
	if state.Turn == Blue{
		val = MinValue(state, alpha, beta, depth-1)		
	}else if state.Turn == Red{
		val = MaxValue(state, alpha, beta, depth-1)		
	}
	return val
}
