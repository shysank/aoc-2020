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

/*

input:
```
BBFFBBFRLL
FFFBBBFRRR
BFFFBBFRRR

```

*/

func ParseBoardingPasses(reader io.Reader) []string {
	var boardingPasses []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		boardingPasses = append(boardingPasses, text)
	}

	return boardingPasses
}

/*

input:
```
abc

a
b
c

```

*/

func ParseCustomsAnswers(reader io.Reader) []answer {
	var answers []answer
	scanner := bufio.NewScanner(reader)
	groupAnswer := answer{yeses: map[rune]int{}}
	var noOfPeople int
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			groupAnswer.noOfPeople = noOfPeople
			answers = append(answers, groupAnswer)

			noOfPeople = 0
			groupAnswer = answer{yeses: map[rune]int{}}
			continue
		}
		noOfPeople++
		for _, c := range text {
			groupAnswer.yeses[c] = groupAnswer.yeses[c] + 1
		}
	}
	groupAnswer.noOfPeople = noOfPeople
	answers = append(answers, groupAnswer)
	return answers
}

/*

input:
```
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
```

*/

func ParseBags(reader io.Reader) (map[bagType][]contains, error) {
	var bags = make(map[bagType][]contains)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, " ")

		// bagType
		var color = ""
		var rem []string
		for i, p := range parts {
			if p == "bags" {
				rem = parts[i+2:]
				break
			}
			color = color + p + " "
		}
		color = strings.Trim(color, " ")

		if rem[0] == "no" && rem[1] == "other" && rem[2] == "bags." {
			bags[bagType(color)] = []contains{}
			continue
		}

		var containsList []contains
		curr := contains{}
		var newBag bool = true
		for _, p := range rem {
			if newBag {
				n, err := strconv.ParseInt(p, 10, 32)
				if err != nil {
					return nil, err
				}
				curr.qty = int(n)
				newBag = false
				continue
			}
			if strings.Contains(p, ".") || strings.Contains(p, ",") {
				curr.color = bagType(strings.Trim(string(curr.color), " "))
				containsList = append(containsList, curr)
				curr = contains{}
				newBag = true
				continue
			}

			curr.color = bagType(string(curr.color) + p + " ")
		}
		bags[bagType(color)] = containsList
	}
	return bags, nil
}

/*

input:
```
nop +0
acc +1
jmp +4
acc +3
```

*/

func ParseBootCode(reader io.Reader) (instructions []instruction, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, " ")
		if len(parts) != 2 {
			return nil, errors.New("Expected instruction of form `nop +0`")
		}

		kind := instructionType(parts[0])
		arg, _ := strconv.ParseInt(parts[1], 10, 32)

		instructions = append(instructions, instruction{kind, int(arg)})
	}

	return instructions, nil
}

/*

input:
```
35
20
15
25
47
40
```

*/

func ParseToInt64Array(reader io.Reader) (nos []int64, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		n, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return nil, err
		}
		nos = append(nos, n)
	}

	return nos, nil
}
