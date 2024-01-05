package game

type Game struct {
	CurrentState State
}

func NewGame() Game {
	state := State{Turn: Red}
	state.Move(N/2, N/2-1)
	state.Move(N/2-1, N/2)
	return Game{CurrentState: state}
}

func MinValue(state State, depth int) State {
	oacs := state.ExpandOpponentActions()
	if state.IsTerminal {
		//time.Sleep(2 * time.Second)
		//state.Print()
		return state
	}
	min := State{Counts: [3]int8{127, 127, 127}}
	var term State
	for idx := range oacs {
		term = MaxValue(oacs[idx], depth+1)
		if min.Counts[Blue] > term.Counts[Blue] {
			min = oacs[idx]
		}
	}
	return min
}

func MaxValue(state State, depth int) State {
	oacs := state.ExpandOpponentActions()
	if state.IsTerminal {
		//time.Sleep(2 * time.Second)
		//state.Print()
		return state
	}
	max := State{Counts: [3]int8{-127, -127, -127}}
	var term State
	for idx := range oacs {
		term = MinValue(oacs[idx], depth+1)
		if max.Counts[Blue] < term.Counts[Blue] {
			max = oacs[idx]
		}
	}
	return max
}
