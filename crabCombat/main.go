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
	game := NewGame(players[0], players[1], true)
	game.play()
}

type Game struct {
	id        int
	playerOne *Player
	playerTwo *Player
	winner    *Player
	round     int
	recursive bool
	history   map[string]bool
}

var numberOfGames = 0
var debug = false

func NewGame(playerOne *Player, playerTwo *Player, recursive bool) *Game {
	numberOfGames++
	game := &Game{id: numberOfGames, playerOne: playerOne, playerTwo: playerTwo, recursive: recursive}
	game.history = make(map[string]bool)
	return game
}

func (g *Game) play() {
	g.printStart()

	for {
		g.round++
		g.printStatus()

		if g.recursive && g.isHistoryRepeating() {
			if debug {
				fmt.Println("History is repeating.")
			}
			g.winner = g.playerOne
			break
		}

		g.saveHistory()

		cardOne, _ := g.playerOne.popCard()
		cardTwo, _ := g.playerTwo.popCard()

		g.checkRoundWinner(cardOne, cardTwo)

		isWinner := g.checkWinner()
		if isWinner {
			break
		}
	}

	g.printEnd()

	if g.id == 1 {
		fmt.Println(fmt.Sprintf("Winner is player number %d with score %d after %d rounds.", g.winner.number, g.winner.getScore(), g.round))
	}
}

func (g *Game) checkRoundWinner(cardOne int, cardTwo int) {
	var winner *Player
	if g.recursive && g.shouldDecideSubGame(cardOne, cardTwo) {
		if debug {
			fmt.Println("Playing a sub-game to determine the winner...")
			fmt.Println()
		}
		subGame := g.createSubGame(cardOne, cardTwo)
		subGame.play()
		if subGame.winner.number == subGame.playerOne.number {
			winner = g.playerOne
		} else {
			winner = g.playerTwo
		}
	} else {
		if cardOne > cardTwo {
			winner = g.playerOne
		} else {
			winner = g.playerTwo
		}
	}

	if winner == g.playerOne {
		g.playerOne.addToBottom(cardOne, cardTwo)
	} else {
		g.playerTwo.addToBottom(cardTwo, cardOne)
	}
	if debug {
		fmt.Printf("Player %d wins round %d of game %d!\n\n", winner.number, g.round, g.id)
	}
}

func (g *Game) printStatus() {
	if debug {
		fmt.Printf("-- Round %d (Game %d) --\n", g.round, g.id)
		fmt.Printf("Player %d's deck: %v\n", g.playerOne.number, g.playerOne.cards)
		fmt.Printf("Player %d's deck: %v\n", g.playerTwo.number, g.playerTwo.cards)
	}
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

func (g *Game) isHistoryRepeating() bool {
	hash := g.playerOne.getCardsHash() + ":" + g.playerTwo.getCardsHash()
	return g.history[hash]
}

func (g *Game) saveHistory() {
	hash := g.playerOne.getCardsHash() + ":" + g.playerTwo.getCardsHash()
	g.history[hash] = true
}

func (g *Game) shouldDecideSubGame(cardOne, cardTwo int) bool {
	return len(g.playerOne.cards) >= cardOne && len(g.playerTwo.cards) >= cardTwo
}

func (g *Game) createSubGame(cardOne, cardTwo int) *Game {
	game := NewGame(&Player{
		number: g.playerOne.number,
		cards:  nil,
	}, &Player{
		number: g.playerTwo.number,
		cards:  nil,
	}, g.recursive)
	for i := 0; i < cardOne; i++ {
		game.playerOne.cards = append(game.playerOne.cards, g.playerOne.cards[i])
	}
	for i := 0; i < cardTwo; i++ {
		game.playerTwo.cards = append(game.playerTwo.cards, g.playerTwo.cards[i])
	}

	return game
}

func (g *Game) printStart() {
	if debug {
		fmt.Printf("=== Game %d ===\n\n", g.id)
	}
}

func (g *Game) printEnd() {
	if debug || g.id == 1 {
		fmt.Println()
		fmt.Println("== Post-game results ==")
		fmt.Printf("Player %d's deck: %v\n", g.playerOne.number, g.playerOne.cards)
		fmt.Printf("Player %d's deck: %v\n", g.playerTwo.number, g.playerTwo.cards)
	}
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
	if debug {
		fmt.Printf("Player %d plays: %d\n", p.number, p.cards[0])
	}
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

func (p *Player) getCardsHash() string {
	var cards []string
	for _, card := range p.cards {
		cards = append(cards, strconv.Itoa(card))
	}
	return strings.Join(cards, ",")
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
