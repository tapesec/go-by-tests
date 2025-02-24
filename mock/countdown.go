package mock

import (
	"fmt"
	"io"
	"iter"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func Countdown(out io.Writer, s Sleeper) {
	for i := range countDownFrom(countdownStart) {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
