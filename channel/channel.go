package channel

import (
	"fmt"
	"time"
)

func attackWithoutChannel(target string) {
	fmt.Println("Throwing ninja stars at", target)
	<-time.After(1 * time.Second)
}

func attackWithChannel(target string, attacked chan bool) {
	fmt.Println("Throwing ninja stars at", target)
	<-time.After(1 * time.Second)
	attacked <- true
}
