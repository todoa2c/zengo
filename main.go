package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var pattern = flag.String("pattern", "", "Search pattern")
var chars = flag.Int("chars", 20, "Characters around pattern")

// foundPos stores position of occurence and line number
type foundPos struct {
	runePos int
	lineNo  int
}

// runeIndex is rune version of strings.Index
func runeIndex(rs, rsep []rune) int {
	if len(rs) < len(rsep) {
		return -1
	}
	for i, _ := range rs {
		matched := true
		p := i
		for j, _ := range rsep {
			if rs[p] != rsep[j] {
				matched = false
				break
			}
			p++
		}
		if matched {
			return i
		}
	}

	return -1
}

// eachRunePos returns position of sep in s
func eachRunePos(s, sep string) []int {
	index := 0
	result := make([]int, 0)
	rs := []rune(s)
	rsep := []rune(sep)
	pos := runeIndex(rs[index:], rsep)
	for pos != -1 {
		result = append(result, index+pos)
		index += pos + 1
		pos = runeIndex(rs[index:], rsep)
	}
	return result
}

func main() {
	flag.Parse()
	if *pattern == "" {
		if flag.NArg() == 0 {
			flag.Usage()
		}
		os.Exit(1)
	}

	runes := make([]rune, 0) //text converted to runes
	lastLen := 0
	foundPosList := make([]foundPos, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for lineNo := 1; scanner.Scan(); lineNo++ {
		line := strings.TrimSpace(scanner.Text()) + " "
		for _, r := range []rune(line) {
			runes = append(runes, r)
		}

		for _, pos := range eachRunePos(line, *pattern) {
			foundPosList = append(foundPosList, foundPos{lastLen + pos, lineNo})
		}

		lastLen = len(runes)
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	runePattern := []rune(*pattern)

	for _, foundPos := range foundPosList {
		pos := foundPos.runePos
		fmt.Printf("%4d\t%s\n", foundPos.lineNo, string(runes[pos-*chars:pos+len(runePattern)+*chars]))
	}
}
