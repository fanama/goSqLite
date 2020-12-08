package console

import (
	"fmt"
	"time"
)

//Log ...
func Log(data string) {
	t := time.Now()
	fmt.Println(t.Local().String(), data)
}

//Timer ...
func Timer(number int) {
	for range time.Tick(time.Second * 5) {
		fmt.Println(number)
		number++
	}
}
