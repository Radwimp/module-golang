package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type MyCaesar struct {
}

var (
	i   int
	c   byte
	key byte
)

func (a MyCaesar) Encode(s string) string {
	res := ""
	key = 3
	for i = 0; i < len(s); i++ {
		c = s[i]
		if c > 64 && c < 91 {
			c += 32
		}
		c += key
		if c > 122 {
			c -= 26
		}
		if c < 97 {
			c += 26
		}
		if c > 96 && c < 123 {
			res += string(c)
		}
	}
	return res
}

func (a MyCaesar) Decode(s string) string {
	res := ""
	key = 3
	for i = 0; i < len(s); i++ {
		c = s[i]
		c -= key
		if c > 122 {
			c -= 26
		}
		if c < 96 {
			c += 26
		}
		if c > 96 && c < 123 {
			res += string(c)
		}
	}
	return res
}

func NewCaesar() Cipher {
	return MyCaesar{}
}
