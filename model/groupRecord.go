package model

type GroupRecord struct {
	GroupID      int
	Team         Team
	NumberOfWin  int
	NumberOfLose int
	NumberOfDraw int
	TotalGoal    int
	TotalScore   int
}
