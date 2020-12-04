package puzzles

import (
	"bufio"
	"errors"
	"fmt"
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
func ParseToIntArray(reader io.Reader) ([]int, error) {
	var inputs []int
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		numText := scanner.Text()
		num, err := strconv.ParseInt(numText, 10, 32)
		if err != nil {
			return nil, err
		}

		inputs = append(inputs, int(num))

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
			return nil, errors.New("Expected 2 tokens. Eg. `1-3` ")
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

/*

input:
```
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
```

*/
func ParseTobogganMap(reader io.Reader) (tobogganMap, error) {
	grid := make(map[tobogganCoordinates]tobogganElement)
	x := -1
	y := -1
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		x = -1
		y++
		for _, c := range text {
			x++
			grid[tobogganCoordinates{x, y}] = tobogganElement(c)
		}
	}

	return tobogganMap{grid: grid, size: tobogganCoordinates{
		x: x + 1,
		y: y + 1,
	}}, nil
}

/*

input:

```
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in

```

*/
func ParsePassports(reader io.Reader) ([]passport, error) {
	var passports []passport
	scanner := bufio.NewScanner(reader)
	p := passport{data: map[string]string{}}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			passports = append(passports, p)
			p = passport{data: map[string]string{}}
			continue
		}

		kvs := strings.Split(text, " ")
		for _, kv := range kvs {
			kvArr := strings.Split(kv, ":")
			if len(kvArr) != 2 {
				return nil, errors.New(fmt.Sprintf("Passport entries should be of the form %q", "key:val"))
			}
			p.data[kvArr[0]] = kvArr[1]
		}
	}
	passports = append(passports, p)

	return passports, nil
}
