package chapter

import "strconv"

func addBinary(a string, b string) string {
	s := ""
	i, j := len(a)-1, len(b)-1
	carry := 0
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		s += strconv.Itoa(sum % 2)
		carry = sum / 2
	}
	if carry != 0 {
		s += strconv.Itoa(carry)
	}
	bs := []byte(s)
	l := len(bs)
	for k := 0; k < l/2; k++ {
		bs[k], bs[l-1-k] = bs[l-1-k], bs[k]
	}
	return string(bs)
}
