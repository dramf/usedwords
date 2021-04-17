package stat

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type WordStat struct {
	uniqueWords map[string]int
	lines       int
	newWords    int
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func InitWordStat(data string, ignoreNumber bool) (*WordStat, error) {
	ws := &WordStat{}
	ws.uniqueWords = map[string]int{}

	lines := strings.Split(data, "\n")
	ws.lines = len(lines)

	for _, l := range lines {
		words := strings.Fields(l)
		for _, w := range words {
			if ignoreNumber && isNumeric(w) {
				continue
			}
			ws.uniqueWords[w]++
		}
	}
	return ws, nil
}

func (ws *WordStat) Words() []string {
	words := []string{}
	if ws != nil && ws.uniqueWords != nil {
		for k, _ := range ws.uniqueWords {
			words = append(words, k)
		}
	}
	return words
}

func (ws *WordStat) NewWords(oldWords []string) []string {
	old := make(map[string]struct{})
	for _, v := range oldWords {
		old[v] = struct{}{}
	}

	newWords := []string{}

	if ws != nil && ws.uniqueWords != nil {
		for k, _ := range ws.uniqueWords {
			_, ok := old[k]
			if !ok {
				newWords = append(newWords, k)
			}
		}
	}
	ws.newWords = len(newWords)
	return newWords
}

func (ws *WordStat) String() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintln("This text contains:"))
	buff.WriteString(fmt.Sprintf("%d lines\n", ws.lines))
	buff.WriteString(fmt.Sprintf("%d uniques words\n", len(ws.uniqueWords)))
	if ws.newWords > 0 {
		buff.WriteString(fmt.Sprintf("%d new words\n", ws.newWords))
	}
	return buff.String()
}
