package cipher

type MyShift struct {
	v int
}

func (a MyShift) Encode(s string) string {
	res := ""
	key = byte(a.v);
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

func (a MyShift) Decode(s string) string {
	res := ""
	key = byte(a.v);
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

func NewShift(v int) Cipher {
	if v > 25 || v < -25 || v == 0 {
		return nil
	}
	return MyShift{v}
}

