package ui

import (
	"fmt"
	"november/internal/model"
	"strings"
)

func Header() {
	fmt.Println(`
  _   _                          _
 | \ | | _____   _____ _ __ ___ | |__   ___ _ __
 |  \| |/ _ \ \ / / _ \ '_ ' _ \| '_ \ / _ \ '__|
 | |\  | (_) \ V /  __/ | | | | | |_) |  __/ |
 |_| \_|\___/ \_/ \___|_| |_| |_|_.__/ \___|_|
                                
`)
	fmt.Println("              Your Stats\n")
}

func Bar(current, max int, width int) string {
	if max == 0 {
		return ""
	}

	filled := (current * width) / max
	empty := width - filled

	return strings.Repeat("█", filled) + strings.Repeat("░", empty)
}

func DisplayUserStats(s model.UserStats) {
	Header()

	max := s.TotalSolved

	fmt.Printf("User: %s\n\n", s.UserName)

	fmt.Println("Solved Problems:\n")
	fmt.Printf("%sEasy	%s  %s  %d\n\n", Green, Bar(s.EasySolved, max, 30), Reset, s.EasySolved)
	fmt.Printf("%sMedium	%s 	%s  %d\n\n", Yellow, Bar(s.MediumSolved, max, 30), Reset, s.MediumSolved)
	fmt.Printf("%sHard	%s  %s  %d\n\n", Red, Bar(s.HardSolved, max, 30), Reset, s.HardSolved)

	fmt.Println("Contest:")
	fmt.Printf("  Rating:           %.2f\n", s.Rating)
	fmt.Printf("  Global Rank:      %d\n", s.GlobalRanking)
	fmt.Printf("  Contests Taken:   %d\n", s.ContestCount)
	fmt.Println()
}
