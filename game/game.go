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

func MinValue(state State, depth int) int8 {
	oacs := state.ExpandOpponentActions()
	if state.IsTerminal {
		return state.Utility()
	}
	var min int8 = math.MaxInt8
	for idx := range oacs {
		term := MaxValue(oacs[idx], depth+1)
		if min > term {
			min = term
		}
	}
	return min
}

func MaxValue(state State, depth int) int8 {
	oacs := state.ExpandOpponentActions()
	if state.IsTerminal {
		return state.Utility()
	}
	var max int8 = math.MinInt8
	for idx := range oacs {
		term := MinValue(oacs[idx], depth+1)
		if max < term {
			max = term
		}
	}
	return max
}
