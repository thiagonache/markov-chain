package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func oldMain() {
	chain := map[string][]string{"hello": {"world", "gophers"}, "world": {"domination", "destruction"}}
	var keys []string
	for k := range chain {
		keys = append(keys, k)
	}
	random := rand.Intn(len(keys))
	word := keys[random]
	var successes []string
	totalRun, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < totalRun; i++ {
		fmt.Printf("%s ", word)
		successes = chain[word]
		if len(successes) == 0 {
			random := rand.Intn(len(keys))
			word = keys[random]
			continue
		}
		randomSuccess := rand.Intn(len(successes))
		word = successes[randomSuccess]
	}
	fmt.Println()
}

func displayChain(chain map[string][]string, n int) {
	prefix := make([]string, 2)
	var words []string
	for i := 0; i < n; i++ {
		key := strings.Join(prefix, " ")
		choices := chain[key]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		copy(prefix, prefix[1:])
		prefix[len(prefix)-1] = next
	}
	fmt.Println(strings.Join(words, " "))
}
func main() {
	stdIn := os.Stdin
	scanner := bufio.NewScanner(stdIn)
	scanner.Split(bufio.ScanWords)

	chain := make(map[string][]string, 2)
	prefix := make([]string, 2)
	for scanner.Scan() {
		input := scanner.Text()
		key := strings.Join(prefix, " ")
		chain[key] = append(chain[key], input)
		copy(prefix, prefix[1:])
		prefix[len(prefix)-1] = input
	}
	displayChain(chain, 20)
}
