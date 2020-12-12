package puzzles

import (
	"io"
	"regexp"
	"strconv"
)

type PassportProcessing struct {
	requiredFields map[string]predicate
	optionalFields map[string]predicate
}

type passport struct {
	data map[string]string
}

type predicate func(string) bool

func NewPassportProcessing() *PassportProcessing {
	return &PassportProcessing{
		requiredFields: map[string]predicate{
			"byr": byrPredicate, "iyr": iyrPredicate, "eyr": eyrPredicate, "hgt": hgtPredicate, "hcl": hclPredicate,
			"ecl": eclPredicate, "pid": pidPredicate,
		},

		optionalFields: map[string]predicate{
			"cid": func(s string) bool {
				return true
			},
		},
	}
}

func (p *PassportProcessing) Puzzle1(reader io.Reader) (Result, error) {
	passports, err := ParsePassports(reader)
	if err != nil {
		return nil, err
	}

	var valPassports int
	for _, passport := range passports {
		requiredFields := copyMap(p.requiredFields)
		optionalFields := copyMap(p.optionalFields)
		if passport.isValid(requiredFields, optionalFields) {
			valPassports++
		}
	}
	return intResult(valPassports), nil
}

func (p *PassportProcessing) Puzzle2(reader io.Reader) (Result, error) {
	passports, err := ParsePassports(reader)
	if err != nil {
		return nil, err
	}

	var valPassports int
	for _, passport := range passports {
		requiredFields := copyMap(p.requiredFields)
		optionalFields := copyMap(p.optionalFields)
		if passport.strictValid(requiredFields, optionalFields) {
			valPassports++
		}
	}
	return intResult(valPassports), nil
}

func (p passport) isValid(requiredFields, optionalFields map[string]predicate) bool {
	for k, _ := range p.data {
		if _, ok := requiredFields[k]; ok {
			delete(requiredFields, k)
		}
	}

	if len(requiredFields) == 0 {
		return true
	}

	return false
}

func (p passport) strictValid(requiredFields, optionalFields map[string]predicate) bool {
	for k, v := range p.data {
		if predicate, ok := requiredFields[k]; ok {
			if predicate(v) {
				delete(requiredFields, k)
			}

		}
	}

	if len(requiredFields) == 0 {
		return true
	}

	return false
}

func copyMap(original map[string]predicate) map[string]predicate {
	copy := make(map[string]predicate)
	for k, v := range original {
		copy[k] = v
	}

	return copy
}

// Predicates

func byrPredicate(val string) bool {
	if len(val) != 4 {
		return false
	}

	n, _ := strconv.ParseInt(val, 10, 32)
	if n >= 1920 && n <= 2002 {
		return true
	}
	return false

}
func iyrPredicate(val string) bool {
	if len(val) != 4 {
		return false
	}

	n, _ := strconv.ParseInt(val, 10, 32)
	if n >= 2010 && n <= 2020 {
		return true
	}
	return false

}

func eyrPredicate(val string) bool {
	if len(val) != 4 {
		return false
	}

	n, _ := strconv.ParseInt(val, 10, 32)
	if n >= 2020 && n <= 2030 {
		return true
	}
	return false

}

func hgtPredicate(val string) bool {
	r := regexp.MustCompile(`([0-9]+)(cm|in)`)
	if !r.MatchString(val) {
		return false
	}
	parts := r.FindStringSubmatch(val)
	n, _ := strconv.ParseInt(parts[1], 10, 32)
	if parts[2] == "cm" {
		if n >= 150 && n <= 193 {
			return true
		}
	}
	if parts[2] == "in" {
		if n >= 59 && n <= 76 {
			return true
		}
	}
	return false
}

func hclPredicate(val string) bool {
	match, _ := regexp.MatchString(`^#([0-9a-f]{6})$`, val)
	return match
}

func eclPredicate(val string) bool {
	match, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, val)
	return match
}

func pidPredicate(val string) bool {
	match, _ := regexp.MatchString(`^([0-9]{9})$`, val)
	return match
}
