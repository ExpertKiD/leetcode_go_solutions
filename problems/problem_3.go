package problems

import (
	"bytes"
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	// return 0 if string is empty
	if len(s) == 0 {
		return 0
	}

	var arr [][]byte

	for i := 0; i < len(s); i++ {
		arr = append(arr, []byte{})

		var strArr = &arr[i]

		for j := i; j < len(s); j++ {
			if bytes.IndexByte(*strArr, s[j]) == -1 {
				*strArr = append(*strArr, s[j])
			} else {
				break
			}
		}
	}

	var longestSubStringIndex = 0
	var longestSubStringlength = 0

	for i := 0; i < len(arr); i++ {

		//fmt.Println(fmt.Printf("%s\n", arr[i]))

		if len(arr[i]) > longestSubStringlength {
			longestSubStringIndex = i
			longestSubStringlength = len(arr[i])
		}

	}

	return len(arr[longestSubStringIndex])
}

func Problem3() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"

	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))

}
