package channel

import (
	"fmt"
	"testing"
	"time"
)

var evilNinjas = []string{
	"Tommy",
	"Johnny",
	"Bobby",
	"Andy",
}

var start = time.Now()

func TestAttackWithoutChannel(t *testing.T) {
	defer func() {
		fmt.Println("duration : ", time.Since(start))
	}()

	for _, evilNinja := range evilNinjas {
		go attackWithoutChannel(evilNinja)
	}

	// tanpa channel, program akan memakan waktu selama waktu yang diatur pada besaran sleep di bawah ini, tak peduli data sudah dieksekusi atau belum
	<-time.After(2 * time.Second)
}

func TestAttackWitChannel(t *testing.T) {
	defer func() {
		fmt.Println("duration : ", time.Since(start))
	}()

	smokeAttack := make(chan bool)

	for _, evilNinja := range evilNinjas {
		go attackWithChannel(evilNinja, smokeAttack)
	}

	fmt.Println(<-smokeAttack)
}
