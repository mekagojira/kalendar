package listevent

import "kalendar/defined"

type Input struct {
	Day   int `json:"day"`
	Month int `json:"month"`
}

type Output struct {
	Events []string `json:"events"`
}

type Ctx = defined.Context[Input, Output]
