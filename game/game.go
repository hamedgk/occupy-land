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

func MinValue(state State, alpha, beta int8) int8 {
	expinfo := state.ExpandOpponentActions()
	if expinfo.IsTerminal {
		return state.Utility()
	}
	var min int8 = math.MaxInt8
	for idx := range expinfo.OpponentActions {
		term := MaxValue(expinfo.OpponentActions[idx], alpha, beta)
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

func MaxValue(state State, alpha, beta int8) int8 {
	expinfo := state.ExpandOpponentActions()
	if expinfo.IsTerminal {
		return state.Utility()
	}
	var max int8 = math.MinInt8
	for idx := range expinfo.OpponentActions {
		term := MinValue(expinfo.OpponentActions[idx], alpha, beta)
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
