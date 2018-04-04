package model

// PointsSorter sorts an array of teams by total points
type PointsSorter []*TeamSeason

func (a PointsSorter) Len() int           { return len(a) }
func (a PointsSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PointsSorter) Less(i, j int) bool { return a[i].TotalPoints() > a[j].TotalPoints() }
