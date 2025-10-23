package main

func hasSameDigits(s string) bool {
	for len(s) > 2 {
		s = reduce(s)
	}

	return s[0] == s[1]
}

func reduce(s string) string {
	n := len(s)
	buf := make([]byte, 0, n-1)
	for i := 0; i < n-1; i++ {
		a := int(s[i] - '0')
		b := int(s[i+1] - '0')
		buf = append(buf, byte((a+b)%10)+'0')
	}
	return string(buf)
}
