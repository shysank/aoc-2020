package puzzles

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

/*

input: ```
input: ```
10
20
30
```

output ```
[10,20,30]
```

*/
func ParseToInt64Array(reader io.Reader) ([]int64, error) {
	var inputs []int64
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		numText := scanner.Text()
		num, err := strconv.ParseInt(numText, 10, 32)
		if err != nil {
			return nil, err
		}

		inputs = append(inputs, num)

	}
	return inputs, nil
}

/*

input: ```
input: ```
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
```

*/
func ParsePasswordRules(reader io.Reader) ([]passwordRules, error) {
	var rules []passwordRules
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		rule := passwordRules{}
		text := scanner.Text()
		tokens := strings.Split(text, " ")
		if len(tokens) != 3 {
			return nil, errors.New("Expected 3 tokens. Eg. `1-3 a: abcde` ")
		}

		// parse range
		rangeParts := strings.Split(tokens[0], "-")
		if len(rangeParts) != 2 {
			return nil, errors.New("Expected 3 tokens. Eg. `1-3` ")
		}

		var err error

		rule.min, err = strconv.ParseInt(rangeParts[0], 10, 32)
		if err != nil {
			return nil, err
		}

		rule.max, err = strconv.ParseInt(rangeParts[1], 10, 32)
		if err != nil {
			return nil, err
		}

		// parse letter
		runes := []rune(strings.Split(tokens[1], ":")[0])
		rule.letter = runes[0]

		//  parse password
		rule.password = tokens[2]

		rules = append(rules, rule)
	}
	
	return rules, nil
}
