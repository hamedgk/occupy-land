package game

import "fmt"

type State struct {
	Board     [N][N]uint8
	Turn       Turn
	IsTerminal bool
	Counts     [3]int8
}

func (state *State) ExpandOpponentActions() []State {
	opponentStates := []State{}
	selfStates := []State{}

	sturn := state.Turn
	oturn := toggleTurn(sturn)

	for x := uint8(0); x < N; x++ {
		for y := uint8(0); y < N; y++ {
			if state.Board[x][y] == oturn {
				state.possibleActions(x, y, &opponentStates, oturn)
			} else if state.Board[x][y] == sturn {
				state.possibleActions(x, y, &selfStates, sturn)
			}
		}
	}
	if len(opponentStates) == 0 && len(selfStates) == 0 {
		state.IsTerminal = true
	} else {
		state.inferTheRest(len(opponentStates), len(selfStates), oturn)
	}
	//empty if terminal
	return opponentStates
}

func (state *State) possibleActions(x, y uint8, states *[]State, turn Turn) {
	switch {
	case x == 0:
		switch {
		case y == 0:
			state.addEmptyNeighbors(x+1, y, states, turn)
			state.addEmptyNeighbors(x, y+1, states, turn)
		case y == N-1:
			state.addEmptyNeighbors(x, y-1, states, turn)
			state.addEmptyNeighbors(x+1, y, states, turn)
		default:
			state.addEmptyNeighbors(x, y-1, states, turn)
			state.addEmptyNeighbors(x, y+1, states, turn)
			state.addEmptyNeighbors(x+1, y, states, turn)
		}
	case x == N-1:
		switch {
		case y == 0:
			state.addEmptyNeighbors(x-1, y, states, turn)
			state.addEmptyNeighbors(x, y+1, states, turn)
		case y == N-1:
			state.addEmptyNeighbors(x-1, y, states, turn)
			state.addEmptyNeighbors(x, y-1, states, turn)
		default:
			state.addEmptyNeighbors(x-1, y, states, turn)
			state.addEmptyNeighbors(x, y-1, states, turn)
			state.addEmptyNeighbors(x, y+1, states, turn)
		}
	default:
		switch {
		case y == 0:
			state.addEmptyNeighbors(x, y+1, states, turn)
			state.addEmptyNeighbors(x-1, y, states, turn)
			state.addEmptyNeighbors(x+1, y, states, turn)
		case y == N-1:
			state.addEmptyNeighbors(x, y-1, states, turn)
			state.addEmptyNeighbors(x-1, y, states, turn)
			state.addEmptyNeighbors(x+1, y, states, turn)
		default:
			state.addEmptyNeighbors(x, y-1, states, turn)
			state.addEmptyNeighbors(x, y+1, states, turn)
			state.addEmptyNeighbors(x-1, y, states, turn)
			state.addEmptyNeighbors(x+1, y, states, turn)
		}
	}
}

func (state *State) addEmptyNeighbors(x, y uint8, states *[]State, sturn Turn) {
	if state.Board[x][y] == None {
		copiedState := State{Board: state.Board, Turn: sturn, Counts: state.Counts}
		copiedState.Board[x][y] = sturn
		copiedState.Counts[sturn]++
		*states = append(*states, copiedState)
	}
}

func (state *State) inferTheRest(ocount, scount int, oturn Turn) {
	sturn := toggleTurn(oturn)
	if ocount == 0 {
		state.IsTerminal = true
		state.Turn = sturn
		for i := uint8(0); i < N; i++ {
			for j := uint8(0); j < N; j++ {
				if state.Board[i][j] == None {
					state.Board[i][j] = sturn
					state.Counts[sturn]++
				} 
			}
		}
	} else if scount == 0 {
		state.IsTerminal = true
		state.Turn = oturn
		for i := uint8(0); i < N; i++ {
			for j := uint8(0); j < N; j++ {
				if state.Board[i][j] == None {
					state.Board[i][j] = oturn
					state.Counts[oturn]++
				}
			}
		}
	}
}

func toggleTurn(stateTurn Turn) Turn {
	if stateTurn == Red {
		return Blue
	} else if stateTurn == Blue {
		return Red
	} else {
		panic("unexpected turn...")
	}
}

func (state *State) Print() {
	fmt.Printf("-----------------------Turn: %v, Blue: %v, Red: %v, Terminal: %v\n", state.Turn, state.Counts[Blue], state.Counts[Red], state.IsTerminal)
	for i := uint8(0); i < N; i++ {
		for j := uint8(0); j < N; j++ {
			if state.Board[i][j] == Blue {
				fmt.Printf(" |B| ")
			} else if state.Board[i][j] == Red {
				fmt.Printf(" |R| ")
			} else {
				fmt.Printf(" | | ")
			}
		}
		fmt.Println()
	}
}

func (state *State) Move(x, y uint8) {
	//buff, turn, isterm, count
	if state.Board[x][y] == None{
		sturn := state.Turn
		oturn := toggleTurn(sturn)
		state.Counts[oturn]++
		state.Board[x][y] = oturn
		state.Turn = oturn
	}
}
