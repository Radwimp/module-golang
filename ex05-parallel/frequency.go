package letter

var (
	m = make(map[rune]int)
	n = make(map[rune]int)
)

//Frequency blablabla
func Frequency(all string) map[rune]int {
	for i := 0; i < len(all); i++ {
		m[rune(all[i])]++
	}
	return m
}

//Frequency blablabla
func freq(all string, c chan map[rune]int) {
	for i := 0; i < len(all); i++ {
		m[rune(all[i])]++
	}
	c <- m
}

//ConcurrentFrequency blahblahblah
func ConcurrentFrequency(a []string) map[rune]int {
	c := make(chan map[rune]int)
	for _, v := range a {
		go freq(v, c)
		n = <-c
	}
	return n
}
