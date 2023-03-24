package shortener

import (
	"regexp"
)

// first letter: len of URL % 52 = map a-ZA-Z
// number of vals mod 52
// number of consonants
// if conflict = add more letters with this approach

const (
	upperCaseA   = 65
	lowerCaseA   = 97
	alphabetSize = 26
	cardinality  = alphabetSize * 2
)

// TODO: to clarify later: All functions close on this var,
// so we're not copying it at all? It's set once in `init`
// and we read from the original
var chars [cardinality * 2]rune

func init() {
	i := 0
	for j := lowerCaseA; j < lowerCaseA+alphabetSize; j++ {
		chars[i] = rune(j)
		i++
	}
	for j := upperCaseA; j < upperCaseA+alphabetSize; j++ {
		chars[i] = rune(j)
		i++
	}
}

// This is not something I'd implement in production,
// just made up, probably mathematically incorrect
// hashing function, with unknown collision probability
// I just want to use go routines for that.
func Shorten(in string) string {
	len := make(chan string, 1)
	val := make(chan string, 1)
	cons := make(chan string, 1)

	go lenBased(in, len)
	go valBased(in, val)
	go consBased(in, cons)

	return <-len + <-val + <-cons
}

// picks character based on length
func lenBased(in string, c chan string) {
	c <- string(chars[(len(in))%cardinality])
}

// picks character based on number of vals
func valBased(in string, c chan string) {
	re := regexp.MustCompile(
		`(?i)[aeiou]`,
	)
	matches := re.FindAllStringIndex(in, -1)

	c <- string(chars[len(matches)%cardinality])
}

// picks character based on number of vals
func consBased(in string, c chan string) {
	re := regexp.MustCompile(
		`(?i)[bcdfghjklmnpqrstvwxys]`,
	)
	matches := re.FindAllStringIndex(in, -1)

	c <- string(chars[len(matches)%cardinality])
}
