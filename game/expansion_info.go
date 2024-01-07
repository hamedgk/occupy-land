package game

type ExpansionInfo struct {
	OpponentActions                                        []State
	IsTerminal                                             bool
	AvailableOpponentActionCount, AvailableSelfActionCount int
	Counts                                                 [3]int8
}
