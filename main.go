package main

import (
	"fmt"
	"time"
)

func main() {
	// menghitung jumlah waktu yang diperlukan untuk attack semua evil ninjas
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Println("duration : ", duration)
	}()

	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Andy"}

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja)
	}

	// main function menunggu goroutine selama 2 detik
	<-time.After(2 * time.Second)

}

func attack(target string) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(1 * time.Second)
}
