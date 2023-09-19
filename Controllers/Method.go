package Controllers

import "fmt"

type Love struct {
	Partner1 string
	Partner2 string
	LovePer  int
}

func (l Love) LoveCalculator() string {
	return fmt.Sprintf("%v Love %v = %v", l.Partner1, l.Partner2, l.LovePer)
}
