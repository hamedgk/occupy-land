package game

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

type State struct {
	Board  [N][N]uint8
	Turn   Turn
	Counts [3]int8
}

func (state *State) ExpandOpponentActions() ExpansionInfo {
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
		return ExpansionInfo{
			OpponentActions:              nil,
			IsTerminal:                   true,
			AvailableOpponentActionCount: 0,
			AvailableSelfActionCount:     0,
		}
	}
	return ExpansionInfo{
		OpponentActions:              opponentStates,
		IsTerminal:                   false,
		AvailableOpponentActionCount: len(opponentStates),
		AvailableSelfActionCount:     len(selfStates),
	}
}

func (state *State) possibleActions(x, y uint8, states *[]State, turn Turn) {
	switch {
	case x == 0:
		switch {
		case y == 0:
			state.applyOneMove(x+1, y, states, turn)
			state.applyOneMove(x, y+1, states, turn)
		case y == N-1:
			state.applyOneMove(x, y-1, states, turn)
			state.applyOneMove(x+1, y, states, turn)
		default:
			state.applyOneMove(x, y-1, states, turn)
			state.applyOneMove(x, y+1, states, turn)
			state.applyOneMove(x+1, y, states, turn)
		}
	case x == N-1:
		switch {
		case y == 0:
			state.applyOneMove(x-1, y, states, turn)
			state.applyOneMove(x, y+1, states, turn)
		case y == N-1:
			state.applyOneMove(x-1, y, states, turn)
			state.applyOneMove(x, y-1, states, turn)
		default:
			state.applyOneMove(x-1, y, states, turn)
			state.applyOneMove(x, y-1, states, turn)
			state.applyOneMove(x, y+1, states, turn)
		}
	default:
		switch {
		case y == 0:
			state.applyOneMove(x, y+1, states, turn)
			state.applyOneMove(x-1, y, states, turn)
			state.applyOneMove(x+1, y, states, turn)
		case y == N-1:
			state.applyOneMove(x, y-1, states, turn)
			state.applyOneMove(x-1, y, states, turn)
			state.applyOneMove(x+1, y, states, turn)
		default:
			state.applyOneMove(x, y-1, states, turn)
			state.applyOneMove(x, y+1, states, turn)
			state.applyOneMove(x-1, y, states, turn)
			state.applyOneMove(x+1, y, states, turn)
		}
	}
}

func (state *State) applyOneMove(x, y uint8, states *[]State, sturn Turn) {
	if state.Board[x][y] == None {
		copiedState := State{Board: state.Board, Turn: sturn, Counts: state.Counts}
		copiedState.Board[x][y] = sturn
		copiedState.Counts[sturn]++
		*states = append(*states, copiedState)
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
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	greenBold := color.New(color.FgGreen, color.Bold)
	red := color.New(color.FgRed, color.Bold)
	color.Cyan("Turn: %v, Blue: %v, Red: %v, \n", state.Turn, state.Counts[Blue], state.Counts[Red])
	for i := uint8(0); i < N; i++ {
		greenBold.Fprintf(w, "\t%d\t", i)
	}
	fmt.Fprintf(w, "\n")
	for i := uint8(0); i < N; i++ {
		for j := uint8(0); j < N; j++ {
			if state.Board[i][j] == Blue {
				red.Fprintf(w, "\tB\t")
			} else if state.Board[i][j] == Red {
				red.Fprintf(w, "\tR\t")
			} else {
				red.Fprintf(w, "\t \t")
			}
		}
		greenBold.Fprintf(w, "\t%d\t", i)
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
}

func (state *State) Move(x, y uint8) {
	if state.Board[x][y] == None {
		sturn := state.Turn
		oturn := toggleTurn(sturn)
		state.Counts[oturn]++
		state.Board[x][y] = oturn
		state.Turn = oturn
	}
}

func (state *State) MoveVoid() {
	sturn := state.Turn
	oturn := toggleTurn(sturn)
	state.Turn = oturn
}

func (state *State) Utility() int8 {
	return state.Counts[Blue]
}
