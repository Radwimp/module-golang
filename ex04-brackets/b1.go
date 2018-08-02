package brackets

//mport "fmt"

func Bracket(s string) (bool, error) {
	a, b, c := true, true, true
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			a = false
			for j := i + 1; j < len(s); j += 2 {
				if s[j] == ')' {
					a = true
					break
				}
			}
		}
		if s[i] == '[' {
			b = false
			for j := i + 1; j < len(s); j += 2 {
				if s[j] == ']' {
					b = true
					break
				}
			}
		}
		if s[i] == '{' {
			c = false
			for j := i + 1; j < len(s); j += 2 {
				if s[j] == '}' {
					c = true
					break
				}
			}
		}
	}
	if a && b && c {
		return true, nil
	}
	return false, nil
}

/*
func main() {
	fmt.Println(Bracket("{[}]"))
}
*/
