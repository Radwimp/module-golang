package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	waitroom [3]int
	clientid int
	cut      bool
	i        int
)

func cutting(clientid int) int {
	cut = true
	fmt.Print("Client ", clientid, " goes to barber\n")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("Client ", clientid, " is done.\n")
	checkwr()
	cut = false
	return clientid
}

func clientcome(clientid int) {
	if cut {
		if waitroom[0] == 0 {
			waitroom[0] = clientid
		} else if waitroom[1] == 0 {
			waitroom[1] = clientid
		} else if waitroom[2] == 0 {
			waitroom[2] = clientid
		}
	} else {
		go cutting(clientid)
	}
	i++
	return
}

func checkwr() {
	if waitroom[0] != 0 {
		tmp := waitroom[0]
		waitroom[0] = waitroom[1]
		waitroom[1] = waitroom[2]
		waitroom[2] = 0
		go cutting(tmp)
	}
	return
}

func main() {
	rnd := time.Duration(rand.Intn(600)) * time.Millisecond
	client := time.Tick(500 * time.Millisecond)
	stop := time.After(10 * time.Second)
	fmt.Println("Barbershop is open")
	for {
		select {
		case <-client:
			clientid++
			fmt.Print("Client ", clientid, " is come\n")
			go clientcome(clientid)
			rnd = time.Duration(rand.Intn(150)+280) * time.Millisecond
			client = time.Tick(rnd)
			fmt.Println("Client spawntime:", rnd)
		case <-stop:
			fmt.Println("Not cuted clients(waitroom):", waitroom)
			fmt.Println("Barbershop is closing")
			return
		default:
			fmt.Println("...........................")
			time.Sleep(200 * time.Millisecond)
		}
	}
	fmt.Println(waitroom)
	return
}
