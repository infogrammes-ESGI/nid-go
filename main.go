package main

import (
	"bufio"
	"fmt"
	"nid-go/cmd/ruleparser"
	"os"
)

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}

func main() {
	fi := bufio.NewReader(os.NewFile(0, "stdin"))

	for {
		var eqn string
		var ok bool

		fmt.Printf("Rule: ")
		if eqn, ok = readline(fi); ok {
			ruleparser.RuleParserParse(&ruleparser.RuleParserLex{S: eqn})
		} else {
			break
		}
	}
}
