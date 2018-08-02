package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(n int, t float64, wg *sync.WaitGroup, pw []bool) {
	fmt.Printf("worker:%d sleep:%.1f\n", n+1, t)
	pw[n] = true
	t *= 1000
	//	fmt.Println(t)
	x := time.Duration(t) * time.Millisecond
	//	fmt.Println(x)
	time.Sleep(x)
	pw[n] = false
	time.Sleep(5 * time.Millisecond)
	if !pw[n] {
		fmt.Printf("worker:%d stopping\n", n+1)
	}
	wg.Done()
	return
}

//Run asdf
func Run(poolSize int) {
	var wg sync.WaitGroup
	var i int
	pw := make([]bool, poolSize)
	reader := bufio.NewScanner(os.Stdin)
	for ; reader.Scan(); i++ {
		a := reader.Text()
		t, _ := strconv.ParseFloat(a[:], 64)
		wg.Add(1)
		for pw[i%len(pw)] {
		}
		if i < len(pw) {
			fmt.Printf("worker:%d spawning\n", i%len(pw)+1)
		}
		go worker(i%len(pw), t, &wg, pw)
	}
	wg.Wait()
	return
}
