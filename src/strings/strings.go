package strings

import "unicode/utf8"

// Explode a string into a slice of strings
// s is the string to explode
// n is the number of strings to return, in other words, the length of the slice
func explode(s string, n int) []string {

	// Get the length of the string in runes
	l := utf8.RuneCountInString(s)

	// If n is less than 0 or greater than the length of the string, set it to the length of the string
	if n < 0 || n > l {
		n = l
	}

	a := make([]string, n)
	// Iterate over the string
	for i:=0; i<n; i++ {
		//utf8.DecodeRuneInString returns the first rune in the string and its size in bytes
		_, size := utf8.DecodeRuneInString(s)
		// Get the first size bytes of the string
		// what does it mean the sintax for the slice? [:size]
		// It means that it will get the first size bytes of the string
		a[i] = s[:size]
		// Get the rest of the string
		s = s[size:]
	}

	if n > 0 {
		a[n-1] = s
	}

	return a
}