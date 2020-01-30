Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:
```
Input: s = "anagram", t = "nagaram"
Output: true
```

Example 2:
```
Input: s = "rat", t = "car"
Output: false
```

Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?

```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	symbols := make(map[rune]int8, len(s))
	for _, v := range s {
		if _, ok := symbols[v]; !ok {
			symbols[v] = 1
		} else {
			symbols[v]++
		}
	}
	
	for _, v := range t {
		if _, ok := symbols[v]; !ok {
			return false
		} else {
			symbols[v]--
		}
	}
	
	for _, v := range symbols {
		if v != 0 {
			return false
		}
	}
	
	return true
}
```
