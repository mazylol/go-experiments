package counter

import (
	"fmt"
	"time"
)

func Counter() {
	go increment()
	fmt.Println("Press the enter key to stop any time")
	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}

func increment() {
	i := 0
	for {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
		i++
	}
}
