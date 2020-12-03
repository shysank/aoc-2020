package puzzles

import "io"

type passwordPhilosophy struct {
}

type passwordRules struct {
	min      int64
	max      int64
	letter   rune
	password string
}

func (p *passwordPhilosophy) Puzzle1(reader io.Reader) (result, error) {
	rules, err := ParsePasswordRules(reader)
	if err != nil {
		return result{}, err
	}

	var validPasswords []string
	for _, rule := range rules {
		if rule.isValid() {
			validPasswords = append(validPasswords, rule.password)
		}
	}

	return result{day2: day2Result{p1: day2p1Result{validPasswords: validPasswords}}}, nil
}

func (p *passwordPhilosophy) Puzzle2(reader io.Reader) (result, error) {

	rules, err := ParsePasswordRules(reader)
	if err != nil {
		return result{}, err
	}

	var validPasswords []string
	for _, rule := range rules {
		if rule.isValidNew() {
			validPasswords = append(validPasswords, rule.password)
		}
	}

	return result{day2: day2Result{p2: day2p2Result{validPasswords: validPasswords}}}, nil
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
	var occurences int
	for i, c := range r.password {
		if i == int(r.min-1) || i == int(r.max-1) {
			if c == r.letter {
				occurences++
			}
		}
	}

	if occurences != 1 {
		return false
	}

	return true
}
