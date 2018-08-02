package downcase

var (
	res string
	c byte
)

func Downcase(s string) (res string, err error) {
	for i := 0; i < len(s); i++ {
		c = s[i]
		if c > 64 && c < 91 {
			c += 32
		}
		res += string(c)
	}
	return res, err
}

