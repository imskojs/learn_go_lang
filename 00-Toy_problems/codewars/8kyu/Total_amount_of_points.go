package kyu8

import (
	"strconv"
	"strings"
)

func Points(games []string) int {
	var point int
	for _, str := range games {
		scores := strings.Split(str, ":")
		x, _ := strconv.ParseInt(scores[0], 10, 64)
		y, _ := strconv.ParseInt(scores[1], 10, 64)
		switch {
		case x > y:
			point += 3
		case x == y:
			point += 1
		}
	}
	return point
}

func Points1(games []string) int {
	points := 0
	for _, g := range games {
		if g[0] > g[2] {
			points += 3
		} else if g[0] == g[2] {
			points += 1
		}
	}
	return points
}
