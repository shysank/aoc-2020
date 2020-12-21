package puzzles

import (
	"io"
	"strings"
)

type AllergenAssessment struct {
}

type food struct {
	ingredients []string
	allergens   []string
}

type foods []food

const (
	occurences = "occurences"
)

func (a AllergenAssessment) Puzzle1(reader io.Reader) (Result, error) {
	fs, _ := ParseFood(reader)
	assigned := fs.assignIngredients()
	count := 0
	for _, f := range fs {
		for _, ing := range f.ingredients {
			if _, ok := assigned[ing]; !ok {
				count++
			}
		}
	}

	return intResult(count), nil
}

func (a AllergenAssessment) Puzzle2(reader io.Reader) (Result, error) {
	fs, _ := ParseFood(reader)
	assigned := fs.assignIngredients()

	allergenToIng := make(map[string]string)
	var sortedAllergens = make([]string, len(assigned))
	for k, v := range assigned {
		allergenToIng[v] = k
		sortedAllergens = insertSortString(sortedAllergens, v)
	}
	var result []string
	for _, a := range sortedAllergens {
		result = append(result, allergenToIng[a])
	}
	return stringArrayResult(result), nil
}

func (fs foods) assignIngredients() map[string]string {
	assigned := make(map[string]string)

	r, _ := fs.reverseIndex()
	for !allAssigned(r) {

		for allergen, allergenIndex := range r {
			for ing, count := range allergenIndex {
				if ing == occurences {
					continue
				}

				if count < allergenIndex[occurences] {
					delete(allergenIndex, ing)
					continue
				}

				if _, ok := assigned[ing]; ok {
					delete(allergenIndex, ing)
					continue
				}

				if len(allergenIndex) == 2 {
					assigned[ing] = allergen
				}
			}
		}

	}
	return assigned
}

func (fs foods) reverseIndex() (map[string]map[string]int, map[string]bool) {
	revLookup := make(map[string]map[string]int)
	allIngredients := make(map[string]bool)
	for _, f := range fs {
		for _, a := range f.allergens {
			for _, i := range f.ingredients {
				allIngredients[i] = true
				if v, ok := revLookup[a]; ok {
					v[i]++
					if v[i] > v[occurences] {
						v[occurences] = v[i]
					}
				} else {
					revLookup[a] = map[string]int{i: 1, occurences: 1}
				}
			}
		}
	}
	return revLookup, allIngredients
}

func allAssigned(revIndex map[string]map[string]int) bool {
	for _, v := range revIndex {
		if len(v) > 1 {
			return false
		}
	}
	return true
}

func insertSortString(arr []string, val string) []string {
	var index int
	for i, s := range arr {
		if strings.Compare(val, s) > 0 && s != "" {
			continue
		}
		index = i
		break
	}
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = val
	return arr
}
