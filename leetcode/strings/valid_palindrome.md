Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:
```
Input: "A man, a plan, a canal: Panama"
Output: true
```

Example 2:
```
Input: "race a car"
Output: false
```

```go
func isPalindrome(s string) bool {
	lowerCaseAbcNum := []byte{97, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	upperCaseAbsNum := []byte{65, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	
	var str []byte
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(lowerCaseAbcNum); j++ {
			if s[i] == lowerCaseAbcNum[j] || s[i] == upperCaseAbsNum[j] {
				str = append(str, lowerCaseAbcNum[j])
				break
			}
		}
	}
	
	l := len(str)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-1-i] {
			return false
		}
	}
	
	return true
}
```
