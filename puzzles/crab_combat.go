package puzzles

import (
	"io"
	"strconv"
)

type CrabCombat struct {
}

type player struct {
	name  string
	cards []int
}

type combatGame struct {
	player1 *player
	player2 *player

	winner *player

	prevRounds []round
}

type round struct {
	deck1, deck2 []int
}

const (
	player1 = "Player 1"
	player2 = "Player 2"
)

func (c CrabCombat) Puzzle1(reader io.Reader) (Result, error) {
	combat, err := ParseGame(reader)
	if err != nil {
		return nil, err
	}

	combat.play()
	score := combat.winner.score()

	return intResult(score), nil
}

func (c CrabCombat) Puzzle2(reader io.Reader) (Result, error) {
	combat, err := ParseGame(reader)
	if err != nil {
		return nil, err
	}

	combat.playRecursive()
	score := combat.winner.score()

	return intResult(score), nil
}

func (g *combatGame) play() {
	for {
		if len(g.player1.cards) == 0 {
			g.winner = g.player2
			break
		}

		if len(g.player2.cards) == 0 {
			g.winner = g.player1
			break
		}

		c1 := g.player1.draw()
		c2 := g.player2.draw()
		if c1 > c2 {
			g.player1.cards = append(g.player1.cards, c1, c2)
		} else {
			g.player2.cards = append(g.player2.cards, c2, c1)
		}
	}
}

func (g *combatGame) playRecursive() {
	for {

		if len(g.player1.cards) == 0 {
			g.winner = g.player2
			break
		}

		if len(g.player2.cards) == 0 {
			g.winner = g.player1
			break
		}

		for _, r := range g.prevRounds {
			if arraysEqual(r.deck1, g.player1.cards) && arraysEqual(r.deck2, g.player2.cards) {
				g.winner = g.player1
				return
			}
		}

		rd := round{deck1: make([]int, len(g.player1.cards)), deck2: make([]int, len(g.player2.cards))}
		copy(rd.deck1, g.player1.cards)
		copy(rd.deck2, g.player2.cards)
		g.prevRounds = append(g.prevRounds, rd)

		c1 := g.player1.draw()
		c2 := g.player2.draw()

		// sub game
		if len(g.player1.cards) >= c1 && len(g.player2.cards) >= c2 {
			subGame := combatGame{player1: &player{name: player1, cards: make([]int, c1)},
				player2: &player{name: player2, cards: make([]int, c2)}}

			copy(subGame.player1.cards, g.player1.cards[:c1])
			copy(subGame.player2.cards, g.player2.cards[:c2])

			subGame.playRecursive()

			if subGame.winner.name == player1 {
				g.player1.cards = append(g.player1.cards, c1, c2)
			} else if subGame.winner.name == player2 {
				g.player2.cards = append(g.player2.cards, c2, c1)
			} else {
				panic("Player not found")
			}
		} else {
			if c1 > c2 {
				g.player1.cards = append(g.player1.cards, c1, c2)
			} else {
				g.player2.cards = append(g.player2.cards, c2, c1)
			}
		}
	}
}

func (p *player) draw() int {
	c := p.cards[0]
	p.cards = p.cards[1:]
	return c
}

func (p *player) score() int {
	m := 1
	score := 0
	for i := len(p.cards) - 1; i >= 0; i-- {
		score += p.cards[i] * m
		m++
	}
	return score
}

func (g *combatGame) toString() string {
	return g.player1.toString() + g.player2.toString()
}

func (p *player) toString() string {
	playerStr := p.name
	for _, c := range p.cards {
		playerStr += strconv.Itoa(c) + ","
	}
	return playerStr
}
