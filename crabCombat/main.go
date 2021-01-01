package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := `Player 1:
24
22
26
6
14
19
27
17
39
34
40
41
23
30
36
11
28
3
10
21
9
50
32
25
8

Player 2:
48
49
47
15
42
44
5
4
13
7
20
43
12
37
29
18
45
16
1
46
38
35
2
33
31`

	players := parseInput(input)
	game := Game{playerOne: players[0], playerTwo: players[1]}
	game.play()
	fmt.Println(fmt.Sprintf("Winner is player number %d with score %d after %d rounds.", game.winner.number, game.winner.getScore(), game.round))
}

type Game struct {
	playerOne *Player
	playerTwo *Player
	winner    *Player
	round     int
}

func (g *Game) play() {
	for {
		g.round++
		g.printStatus()
		cardOne, _ := g.playerOne.popCard()
		cardTwo, _ := g.playerTwo.popCard()

		if cardOne > cardTwo {
			g.playerOne.addToBottom(cardOne, cardTwo)
		} else {
			g.playerTwo.addToBottom(cardTwo, cardOne)
		}

		isWinner := g.checkWinner()
		if isWinner {
			break
		}
	}
}

func (g *Game) printStatus() {
	fmt.Printf("-- Round %d --\n", g.round)
	fmt.Printf("Player %d's deck: %v\n", g.playerOne.number, g.playerOne.cards)
	fmt.Printf("Player %d's deck: %v\n", g.playerTwo.number, g.playerTwo.cards)
}

func (g *Game) checkWinner() bool {
	if len(g.playerOne.cards) == 0 {
		g.winner = g.playerTwo
		return true
	}
	if len(g.playerTwo.cards) == 0 {
		g.winner = g.playerOne
		return true
	}
	return false
}

type Player struct {
	number int
	cards  []int
}

func (p *Player) popCard() (int, bool) {
	if len(p.cards) < 1 {
		return 0, false
	}
	defer func() { p.cards = p.cards[1:] }()
	return p.cards[0], true
}

func (p *Player) String() string {
	return "Player " + strconv.Itoa(p.number) + ": " + fmt.Sprint(p.cards)
}

func (p *Player) addToBottom(cards ...int) {
	p.cards = append(p.cards, cards...)
}

func (p *Player) getScore() int {
	score := 0
	for i := len(p.cards); i > 0; i-- {
		fmt.Printf("+ %d * %d\n", i, p.cards[len(p.cards)-i])
		score += i * p.cards[len(p.cards)-i]
	}
	return score
}

func parseInput(input string) []*Player {
	var players []*Player
	for _, playerRaw := range strings.Split(input, "\n\n") {
		rows := strings.Split(playerRaw, "\n")
		number, _ := strconv.Atoi(rows[0][7:8])
		var cards []int
		for _, cardRaw := range rows[1:] {
			card, _ := strconv.Atoi(cardRaw)
			cards = append(cards, card)
		}
		players = append(players, &Player{number, cards})
	}
	return players
}
