package puzzles

import (
	"fmt"
	"io"
)

type passwordPhilosophy struct {
}

type passwordRules struct {
	min      int64
	max      int64
	letter   rune
	password string
}

type validPasswords []string

func (p *passwordPhilosophy) Puzzle1(reader io.Reader) (Result, error) {
	rules, err := ParsePasswordRules(reader)
	if err != nil {
		return nil, err
	}

	var vpwds validPasswords
	for _, rule := range rules {
		if rule.isValid() {
			vpwds = append(vpwds, rule.password)
		}
	}

	return vpwds, nil
}

func (p *passwordPhilosophy) Puzzle2(reader io.Reader) (Result, error) {

	rules, err := ParsePasswordRules(reader)
	if err != nil {
		return nil, err
	}

	var vpwds validPasswords
	for _, rule := range rules {
		if rule.isValidNew() {
			vpwds = append(vpwds, rule.password)
		}
	}

	return vpwds, nil
}

func (r passwordRules) isValid() bool {
	var occurrences int64
	for _, c := range r.password {
		if c == r.letter {
			occurrences++
		}
	}

	if occurrences >= r.min && occurrences <= r.max {
		return true
	}

	return false
}

func (r passwordRules) isValidNew() bool {
	var occurrences int
	for i, c := range r.password {
		if i == int(r.min-1) || i == int(r.max-1) {
			if c == r.letter {
				occurrences++
			}
		}
	}

	if occurrences != 1 {
		return false
	}

	return true
}

func (v validPasswords) Value() string {
	return fmt.Sprintf("%d", len(v))
}
