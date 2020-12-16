package puzzles

import (
	"io"
	"strings"
)

type TicketTranslation struct {
}

type ticketRule struct {
	start, end int
}

type ticket struct {
	vals []int
}

func (t TicketTranslation) Puzzle1(reader io.Reader) (Result, error) {
	rulesMap, _, nearbyTickets, err := ParseTickets(reader)
	if err != nil {
		return nil, err
	}

	var errorRateSum int
	for _, t := range nearbyTickets {
		if errorRate, valid := t.errorRate(rulesMap); !valid {
			errorRateSum += errorRate
		}
	}

	return intResult(errorRateSum), nil

}

func (t TicketTranslation) Puzzle2(reader io.Reader) (Result, error) {
	rulesMap, yourTicket, nearbyTickets, err := ParseTickets(reader)
	if err != nil {
		return nil, err
	}

	var validTickets []ticket
	for _, t := range nearbyTickets {
		if _, valid := t.errorRate(rulesMap); valid {
			validTickets = append(validTickets, t)
		}
	}

	var fieldsColArranged [][]int
	for i := 0; i < len(yourTicket.vals); i++ {
		var colVals = make([]int, len(validTickets))
		for j, v := range validTickets {
			colVals[j] = v.vals[i]
		}
		fieldsColArranged = append(fieldsColArranged, colVals)
	}

	possibilities := make(map[int][]string)
	for i := 0; i < len(yourTicket.vals); i++ {
		updatePossibilities(fieldsColArranged[i], i, rulesMap, possibilities)
	}

	assigned := make(map[int]string)
	for len(assigned) < len(yourTicket.vals) {
		for i := 0; i < len(yourTicket.vals); i++ {
			possibilitiesCopy := clonePossibilities(possibilities)
			tryAssign(i, possibilitiesCopy[i], assigned)
		}
	}

	product := 1
	for i, v := range yourTicket.vals {
		if strings.Contains(assigned[i], "departure") {
			product *= v
		}
	}

	return intResult(product), nil
}

func (tr ticketRule) apply(val int) bool {
	return val >= tr.start && val <= tr.end
}

func (t ticket) errorRate(rulesMap map[string][]ticketRule) (int, bool) {
	var errorRate int
	var ticketValid = true
	for _, v := range t.vals {
		valid := false
		for _, rules := range rulesMap {
			for _, r := range rules {
				if r.apply(v) {
					valid = true
					break
				}
			}
		}
		if !valid {
			ticketValid = false
			errorRate += v
		}

	}
	return errorRate, ticketValid
}

func tryAssign(col int, possibilities []string, assigned map[int]string) {
	var notAssigned []string
	for _, p := range possibilities {
		if isAssigned(p, assigned) {
			continue
		}
		notAssigned = append(notAssigned, p)
	}
	if len(notAssigned) == 1 {
		assigned[col] = notAssigned[0]
	}
}

func isAssigned(s string, assigned map[int]string) bool {
	for _, v := range assigned {
		if s == v {
			return true
		}
	}
	return false
}

func updatePossibilities(fields []int, col int, rulesMap map[string][]ticketRule, possibilities map[int][]string) {
	possibleMap := make(map[string]bool)
	impossibilities := make(map[string]bool)
	for _, f := range fields {

		for k, rules := range rulesMap {
			valid := false
			for _, r := range rules {
				if r.apply(f) {
					valid = true
					break
				}
			}
			if valid {
				possibleMap[k] = true
			} else {
				impossibilities[k] = true
			}
		}
	}
	var possibleFields []string
	for k, _ := range possibleMap {
		if impossibilities[k] {
			continue
		}
		possibleFields = append(possibleFields, k)
	}
	possibilities[col] = possibleFields
}

func clonePossibilities(p map[int][]string) map[int][]string {
	o := make(map[int][]string)
	for k, v := range p {
		var newV []string = make([]string, len(v))
		copy(newV, v)
		o[k] = newV
	}
	return o
}
