package main

import (
	"fmt"
	"strings"
)

type SplitResult struct {
	lines    []string
	hasExtra bool
}

func SplitText(text string, breaks []uint) SplitResult {
	// split words by whitespace
	words := strings.Split(text, " ")

	lines := make([]string, len(breaks))
	prevBreak := 0

	// populate lines and update previous break for each iteration
	for i, br := range breaks {
		start := prevBreak
		stop := prevBreak + int(br)

		// return early if exceeded breaks bounds
		if max := cap(words); stop > max {
			line := words[start:max]

			lines[i] = strings.Join(line, " ")
			prevBreak = max
			break
		}

		line := words[start:stop]

		lines[i] = strings.Join(line, " ")
		prevBreak += int(br)
	}

	// check if extra text exists
	hasExtra := false

	if len(words[prevBreak:]) > 0 {
		hasExtra = true
	}

	return SplitResult{lines, hasExtra}
}

func main() {
	sentence := "hello what is the name of this app I want to get it online for my tests"

	result := SplitText(sentence, []uint{2, 3, 4, 1})

	for _, text := range result.lines {
		fmt.Println(text)
	}
	fmt.Println("has extra:", result.hasExtra)
}
