package markov

import (
	"bufio"
	"io"
	"math/rand"
)

type Chain map[string][]string

func BuildChain(r io.Reader) Chain {
	result := Chain{}
	lastWord := ""
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		result[lastWord] = append(result[lastWord], word)
		lastWord = word
	}
	return result
}

func Generate(c Chain, howMany int) string {
	curWord := ""
	words := ""
	for i := 0; i < howMany; i++ {
		choices := c[curWord]
		if len(choices) == 0 {
			curWord = ""
			continue
		}
		index := rand.Intn(len(choices))
		word := choices[index]
		words += word + " "
		curWord = word
	}
	return words
}