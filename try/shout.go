package try

import (
	"fmt"
)

var newshout string

func NewShout(shout string, times int) string {

	fmt.Println(shout)
	fmt.Println(times)

	for i := times; i > 0; i-- {
		newshout = newshout + shout
		fmt.Println(i)
		fmt.Println(newshout)
	}
	return newshout
}
