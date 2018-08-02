package cipher

type MyVigenere struct {
	str string
}

func (a MyVigenere) Encode(s string) string {
	res := ""
	j := 0
	for i = 0; i < len(s); i++ {
		c = s[i]
    if c > 64 && c < 91 {
      c += 32
		}
		key = a.str[j % len(a.str)] - 'a'
		c += key
		if c > 122 {
			c -= 26
		}
		if c < 97 {
			c += 26
		}
		if c > 96 && c < 123 {
			res += string(c)
			j++
		}
	}
	return res
}

func (a MyVigenere) Decode(s string) string {
	res := ""
	for i = 0; i < len(s); i++ {
		c = s[i]
		c -= a.str[i % len(a.str)] - 'a'
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

func NewVigenere(str string) Cipher {
	flag := false
	for i = 0; i < len(str); i++ {
		c = str[i]
		if c < 'a' || c > 'z' {
			flag = false
			break
		}
		if c == 'a'{
			continue
		}
		flag = true
	}
	if flag {
		return MyVigenere{str}
	}
	return nil
}

