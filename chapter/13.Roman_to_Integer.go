package chapter

// Roman numerals are usually written largest to smallest from left to right
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i+1 < len(s) && s[i+1] == 'V' {
				sum += 4
				i++
			} else if i+1 < len(s) && s[i+1] == 'X' {
				sum += 9
				i++
			} else {
				sum += romanMap[s[i]]
			}
		case 'X':
			if i+1 < len(s) && s[i+1] == 'L' {
				sum += 40
				i++
			} else if i+1 < len(s) && s[i+1] == 'C' {
				sum += 90
				i++
			} else {
				sum += romanMap[s[i]]
			}
		case 'C':
			if i+1 < len(s) && s[i+1] == 'D' {
				sum += 400
				i++
			} else if i+1 < len(s) && s[i+1] == 'M' {
				sum += 900
				i++
			} else {
				sum += romanMap[s[i]]
			}
		default:
			sum += romanMap[s[i]]
		}
	}
	return sum
}
