package chapter

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for !isValidChar(s[left]) && left < right {
			left++
		}
		for !isValidChar(s[right]) && left < right {
			right--
		}
		if left < right {
			if !caseInsensitiveCompare(s[left], s[right]) {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isValidChar(c byte) bool {
	if (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') {
		return true
	}
	return false
}

func caseInsensitiveCompare(a, b byte) bool {
	if a == b {
		return true
	}
	if !((a >= '0' && a <= '9') || (b >= '0' && b <= '9')) {
		if int(b)-int(a) == 32 || int(b)-int(a) == -32 {
			return true
		}
	}

	return false
}
