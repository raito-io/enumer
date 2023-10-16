package main

import "fmt"

type TitleCaseValue int

const (
	TitleCaseValueOne TitleCaseValue = iota
	titlecasevaluetwo
	titlecasevaluethree
)

func main() {
	ck(TitleCaseValueOne, "titleCaseValueOne")
	ck(titlecasevaluetwo, "titlecasevaluetwo")
	ck(titlecasevaluethree, "titlecasevaluethree")
	ck(-127, "TitleCaseValue(-127)")
	ck(127, "TitleCaseValue(127)")
}

func ck(value TitleCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_title.go: " + str + " vs " + value.String())
	}
}
