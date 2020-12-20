package puzzles

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type MonsterMessages struct {
}

type messageRule struct {
	id        int
	val       string
	subRules1 []int
	subRules2 []int
}

func (m MonsterMessages) Puzzle1(reader io.Reader) (Result, error) {
	rulesMap, words := ParseMessageRules(reader)

	expr := fmt.Sprintf("^%s$", compile(rulesMap, rulesMap[0]))

	r := regexp.MustCompile(expr)

	matches := 0
	for _, w := range words {
		match := r.MatchString(w)
		if match {
			matches++
		}

	}

	return intResult(matches), nil
}

func (m MonsterMessages) Puzzle2(reader io.Reader) (Result, error) {
	rulesMap, words := ParseMessageRules(reader)

	expr := fmt.Sprintf("^%s$", compileLoop(rulesMap, rulesMap[0]))

	r := regexp.MustCompile(expr)

	matches := 0
	for _, w := range words {
		match := r.MatchString(w)
		if match {
			matches++
		}

	}
	return intResult(matches), nil
}

func compile(rulesMap map[int]messageRule, rule messageRule) string {
	if rule.val != "" {
		return rule.val
	}

	subRule1 := ""
	for _, r := range rule.subRules1 {
		subRule1 += compile(rulesMap, rulesMap[r])
	}

	subRule2 := ""
	for _, r := range rule.subRules2 {
		subRule2 += compile(rulesMap, rulesMap[r])
	}

	var subRule string

	if subRule2 == "" {
		subRule = subRule1
	} else {
		subRule = fmt.Sprintf("(%s|%s)", subRule1, subRule2)

	}

	return subRule
}

func compileLoop(rulesMap map[int]messageRule, rule messageRule) string {
	if rule.val != "" {
		return rule.val
	}

	subRule1 := ""
	// we need this for rule 11
	var subRule1Parts []string
	for _, r := range rule.subRules1 {
		compiledResult := compileLoop(rulesMap, rulesMap[r])
		if rule.id == 11 {
			subRule1Parts = append(subRule1Parts, compiledResult)
		}
		subRule1 += compiledResult
	}

	subRule2 := ""
	for _, r := range rule.subRules2 {
		subRule2 += compileLoop(rulesMap, rulesMap[r])
	}

	var subRule string

	if subRule2 == "" {
		subRule = subRule1
	} else {
		subRule = fmt.Sprintf("(%s|%s)", subRule1, subRule2)

	}

	if rule.id == 8 {
		subRule = fmt.Sprintf("(%s+)", subRule)
	}

	if rule.id == 11 {
		var subRulesWithN string
		// hack to check the same number of occurrences in two sub patterns
		for n := 1; n < 10; n++ {
			subRulesWithN += fmt.Sprintf("(%s{%s}%s{%s})", subRule1Parts[0], strconv.Itoa(n), subRule1Parts[1], strconv.Itoa(n))
			subRulesWithN += "|"
		}

		subRulesWithN = strings.Trim(subRulesWithN, "|")
		subRule = fmt.Sprintf("(%s)", subRulesWithN)
	}

	return subRule

}
