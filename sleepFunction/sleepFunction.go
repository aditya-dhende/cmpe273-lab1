package sleepFunction

import (
	"fmt"
	"time"
)

func Sleep_UsingTimeAfter(dur int64) int {

	t1 := time.Now()
	c := <-time.After(time.Second * time.Duration(dur))
	t2 := time.Now()

	//var time_diff int
	time_diff := t2.Second() - t1.Second()
	fmt.Println(c)
	fmt.Println(time_diff)
	return time_diff

}

func main() {
	Sleep_UsingTimeAfter(10)
}
