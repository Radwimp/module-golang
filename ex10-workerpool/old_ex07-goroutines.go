package goroutines

import (
	"sync"
)

func func1(input chan string, ch chan string, flag *sync.Mutex) {
	a := "(" + <-input + ")"
	ch <- a
	flag.Unlock()
}

func func2(ch chan string, flag *sync.Mutex) {
	flag.Lock()
	close(ch)
}

//Process blah-blah-blah
func Process(input chan string) chan string {
	flag := &sync.Mutex{}
	flag.Lock()
	ch := make(chan string)
	go func1(input, ch, flag)
	go func2(ch, flag)
	return ch
}
