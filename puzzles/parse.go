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
```
10
20
30
```
*/
func ParseToIntArray(reader io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(reader)
	return parseToIntArray(scanner)
}

func parseToIntArray(scanner *bufio.Scanner) ([]int, error) {
	var inputs []int
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
```
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
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
func ParseLayout(reader io.Reader) (*layout, error) {
	scanner := bufio.NewScanner(reader)
	return parseLayout(scanner), nil
}

func parseLayout(scanner *bufio.Scanner) *layout {
	grid := make(map[coordinates]string)
	x := -1
	y := -1
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		x = -1
		y++
		for _, c := range text {
			x++
			grid[coordinates{x, y}] = string(c)
		}
	}

	return &layout{grid: grid, size: coordinates{
		x: x + 1,
		y: y + 1,
	}}
}

func ParseLayoutFromString(text string) *layout {
	scanner := bufio.NewScanner(strings.NewReader(text))
	return parseLayout(scanner)
}

/*
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
```
BBFFBBFRLL
FFFBBBFRRR
BFFFBBFRRR
```
*/

func ParseToStringArray(reader io.Reader) []string {
	var result []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		result = append(result, text)
	}

	return result
}

/*
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

/*
```
F10
N3
F7
R90
F11
```
*/

func ParseNavInstructions(reader io.Reader) (ins []navInstruction, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		action := text[0]
		val := text[1:]
		n, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return nil, err
		}
		ins = append(ins, navInstruction{action: string(action), value: int(n)})
	}

	return ins, nil
}

/*
```
939
7,13,x,x,59,x,31,19
```
*/

func ParseNotes(reader io.Reader) (n *notes, err error) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	timestamp, err := strconv.ParseInt(scanner.Text(), 10, 32)
	if err != nil {
		return nil, err
	}

	n = new(notes)
	n.earliestTime = int(timestamp)

	scanner.Scan()
	busIDs := strings.Split(scanner.Text(), ",")
	for _, id := range busIDs {
		if id == "x" {
			n.busIDs = append(n.busIDs, -1)
		} else {
			busId, err := strconv.ParseInt(id, 10, 32)
			if err != nil {
				return nil, err
			}
			n.busIDs = append(n.busIDs, int(busId))
		}
	}

	return n, nil
}

/*
```
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
```
*/

func ParseInitializationProgram(reader io.Reader) (ops bitOperations, err error) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()

		parts := strings.Split(text, "=")
		if len(parts) != 2 {
			return nil, errors.New("Shoule be of form mask = val or mem[x] = val")
		}

		if strings.Contains(parts[0], "mask") {
			maskOp := maskOperation{val: map[int]int{}}
			mask := strings.Trim(parts[1], " ")
			var j = 0
			for i := len(mask) - 1; i >= 0; i-- {
				if string(mask[i]) != "X" {
					n, err := strconv.ParseInt(string(mask[i]), 10, 32)
					if err != nil {
						return nil, err
					}
					maskOp.val[j] = int(n)
				} else {
					maskOp.val[j] = -1
				}
				j++
			}
			ops = append(ops, bitOperation{maskOp: &maskOp})
		} else {
			memOp := memOperation{}

			memLocParts := strings.Split(parts[0], "[")
			memLoc := strings.Trim(memLocParts[1], "] ")
			loc, err := strconv.ParseInt(memLoc, 10, 32)
			if err != nil {
				return nil, err
			}

			val := strings.Trim(parts[1], " ")
			v, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return nil, err
			}
			memOp.loc = loc
			memOp.val = v
			ops = append(ops, bitOperation{memOp: &memOp})
		}

	}

	return ops, nil
}

/*
```
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
```
*/

func ParseTickets(reader io.Reader) (rulesMap map[string][]ticketRule, yourTicket ticket, nearbyTickets []ticket, err error) {
	scanner := bufio.NewScanner(reader)

	// parse rules
	rulesMap = make(map[string][]ticketRule)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		parts := strings.Split(text, ": ")
		ruleName := parts[0]

		rangeParts := strings.Split(parts[1], "or")
		for _, rp := range rangeParts {
			rp = strings.Trim(rp, " ")
			p := strings.Split(rp, "-")
			start, err := strconv.ParseInt(p[0], 10, 32)
			if err != nil {
				return nil, ticket{}, nil, err
			}
			end, err := strconv.ParseInt(p[1], 10, 32)
			if err != nil {
				return nil, ticket{}, nil, err
			}

			if _, ok := rulesMap[ruleName]; !ok {
				rulesMap[ruleName] = []ticketRule{{int(start), int(end)}}
			} else {
				rulesMap[ruleName] = append(rulesMap[ruleName], ticketRule{int(start), int(end)})
			}

		}

	}

	// parse your ticket
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		if text == "your ticket:" {
			yourTicket = ticket{}
			continue
		}

		valsParts := strings.Split(text, ",")
		for _, valStr := range valsParts {
			val, err := strconv.ParseInt(valStr, 10, 32)
			if err != nil {
				return nil, ticket{}, nil, err
			}
			yourTicket.vals = append(yourTicket.vals, int(val))

		}

	}

	// parse nearby tickets
	var curr ticket
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		if text == "nearby tickets:" {
			continue
		}

		curr = ticket{}
		valsParts := strings.Split(text, ",")
		for _, valStr := range valsParts {
			val, err := strconv.ParseInt(valStr, 10, 32)
			if err != nil {
				return nil, ticket{}, nil, err
			}
			curr.vals = append(curr.vals, int(val))

		}
		nearbyTickets = append(nearbyTickets, curr)

	}

	return rulesMap, yourTicket, nearbyTickets, nil
}

func ParseMessageRules(reader io.Reader) (map[int]messageRule, []string) {
	scanner := bufio.NewScanner(reader)

	rulesMap := make(map[int]messageRule)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		m := messageRule{}
		parts := strings.Split(text, ":")

		id, _ := strconv.ParseInt(parts[0], 10, 32)
		m.id = int(id)

		subrules := parts[1]
		subruleParts := strings.Split(subrules, "|")
		subruleParts1 := strings.Split(strings.Trim(subruleParts[0], " "), " ")

		for _, s := range subruleParts1 {
			sid, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				m.val = strings.Trim(s, "\"")
			} else {
				m.subRules1 = append(m.subRules1, int(sid))
			}
		}

		if len(subruleParts) == 2 {
			subruleParts2 := strings.Split(strings.Trim(subruleParts[1], " "), " ")
			for _, s := range subruleParts2 {
				sid, _ := strconv.ParseInt(s, 10, 32)
				m.subRules2 = append(m.subRules2, int(sid))
			}
		}

		rulesMap[int(id)] = m
	}

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return rulesMap, words
}

func ParseTiles(reader io.Reader) (tiles, error) {

	tilesMap := make(map[int]*tile)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		tileText := scanner.Text()
		tilesParts := strings.Split(tileText, "Tile ")
		tileNoStr := strings.Trim(tilesParts[1], ":")
		tileNo, err := strconv.ParseInt(tileNoStr, 10, 32)
		if err != nil {
			return nil, err
		}
		l := parseLayout(scanner)
		tilesMap[int(tileNo)] = &tile{id: int(tileNo), layout: l}
	}
	return tilesMap, nil
}

/*
```
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
```
*/

func ParseFood(reader io.Reader) (foods, error) {
	var foods []food
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, "(contains ")
		ingredients := strings.Split(strings.Trim(parts[0], " "), " ")
		allergensStr := strings.Trim(parts[1], ")")
		allergens := strings.Split(allergensStr, ", ")
		foods = append(foods, food{ingredients: ingredients, allergens: allergens})
	}

	return foods, nil
}

/*
```
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
```
*/

func ParseGame(reader io.Reader) (*combatGame, error) {
	var combat = combatGame{}
	scanner := bufio.NewScanner(reader)

	var player1 player
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		if text == "Player 1:" {
			player1 = player{name: "Player 1"}
			continue
		}
		num, err := strconv.ParseInt(text, 10, 32)
		if err != nil {
			return nil, err
		}
		player1.cards = append(player1.cards, int(num))

	}

	var player2 player
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		if text == "Player 2:" {
			player2 = player{name: "Player 2"}
			continue
		}
		num, err := strconv.ParseInt(text, 10, 32)
		if err != nil {
			return nil, err
		}
		player2.cards = append(player2.cards, int(num))
	}

	combat.player1 = &player1
	combat.player2 = &player2

	return &combat, nil
}

/*
```
sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
```
*/

func ParseTileDirections(reader io.Reader) []tilePath {
	var tiles []tilePath
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()

		t := tilePath{dirs: []string{}}

		dirs := []string{dirEast, dirWest, dirNorthEast, dirNorthWest, dirSouthEast, dirSouthWest}
		var curr string
		for _, d := range text {
			curr += string(d)
			if isPresent(dirs, curr) {
				t.dirs = append(t.dirs, curr)
				curr = ""
			}
		}
		tiles = append(tiles, t)
	}
	return tiles
}

func isPresent(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
