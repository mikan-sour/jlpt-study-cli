package utils

import (
	"fmt"

	"jlpt/src/core/models"
)

func MapLevelToEnum(lvl string) (*models.JLPTLevel, error) {
	var res models.JLPTLevel
	switch lvl {
	case "JLPT N-1":
		res = 1
	case "JLPT N-2":
		res = 2
	case "JLPT N-3":
		res = 3
	case "JLPT N-4":
		res = 4
	case "JLPT N-5":
		res = 5
	default:
		return nil, fmt.Errorf("level provided is not a valid JLPT Level")
	}

	return &res, nil
}

func MapEnumToLevel(lvl models.JLPTLevel) string {
	switch lvl {
	case 1:
		return "JLPT N-1"
	case 2:
		return "JLPT N-2"
	case 3:
		return "JLPT N-3"
	case 4:
		return "JLPT N-4"
	case 5:
		return "JLPT N-5"
	}
	fmt.Printf("received %d...kinda strange", lvl)
	return ""
}
