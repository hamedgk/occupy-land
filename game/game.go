package game

import "math"

type Game struct {
	CurrentState State
}

func NewGame() Game {
	state := State{Turn: Red}
	state.Move(N/2, N/2-1)
	state.Move(N/2-1, N/2)
	return Game{CurrentState: state}
}

func MinValue(state State) int8 {
	oacs, isterm := state.ExpandOpponentActions()
	if isterm{
		return state.Utility()
	}
	var min int8 = math.MaxInt8
	for idx := range oacs {
		term := MaxValue(oacs[idx])
		if min > term {
			min = term
		}
	}
	return min
}

func MaxValue(state State) int8 {
	oacs, isterm := state.ExpandOpponentActions()
	if isterm {
		return state.Utility()
	}
	var max int8 = math.MinInt8
	for idx := range oacs {
		term := MinValue(oacs[idx])
		if max < term {
			max = term
		}
	}
	return max
}
