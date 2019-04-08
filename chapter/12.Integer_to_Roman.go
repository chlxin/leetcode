package chapter

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// notes: 实际上这道题可以用循环来写，会让代码很简洁，但是这道题实在太没意思了，不想花功夫去整理了，有时间再整理
func intToRoman(num int) string {
	romanMap := map[int]byte{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}
	res := make([]byte, 0)
	// 需要计算千，百，十，个位
	thousand := num / 1000
	if thousand >= 0 {
		for i := 0; i < thousand; i++ {
			res = append(res, romanMap[1000])
		}
	}
	hundred := num / 100 % 10
	switch hundred {
	case 0: //nothing
	case 9:
		res = append(res, []byte{romanMap[100], romanMap[1000]}...)
	case 4:
		res = append(res, []byte{romanMap[100], romanMap[500]}...)
	default:
		if hundred >= 5 {
			hundred = hundred - 5
			res = append(res, romanMap[500])
		}
		for i := 0; i < hundred; i++ {
			res = append(res, romanMap[100])
		}
	}

	ten := num / 10 % 10
	switch ten {
	case 0: //nothing
	case 9:
		res = append(res, []byte{romanMap[10], romanMap[100]}...)
	case 4:
		res = append(res, []byte{romanMap[10], romanMap[50]}...)
	default:
		if ten >= 5 {
			ten = ten - 5
			res = append(res, romanMap[50])
		}
		for i := 0; i < ten; i++ {
			res = append(res, romanMap[10])
		}
	}

	one := num % 10
	switch one {
	case 0: //nothing
	case 9:
		res = append(res, []byte{romanMap[1], romanMap[10]}...)
	case 4:
		res = append(res, []byte{romanMap[1], romanMap[5]}...)
	default:
		if one >= 5 {
			one = one - 5
			res = append(res, romanMap[5])
		}
		for i := 0; i < one; i++ {
			res = append(res, romanMap[1])
		}
	}

	return string(res)
}
