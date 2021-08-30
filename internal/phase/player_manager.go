package phase

import (
	"container/list"

	"github.com/jonathan-buttner/game-framework/internal/player"
)

type PlayerManager struct {
	players              *list.List
	currentPlayerElement *list.Element
}

func NewPlayerManager(players []*player.Player) *PlayerManager {
	playerList := list.New()
	for _, p := range players {
		playerList.PushBack(p)
	}
	return &PlayerManager{playerList, playerList.Front()}
}

type PlayersAction func(player *player.Player) error

func (p *PlayerManager) ExecuteForPlayers(action PlayersAction) error {
	for playerElement := p.players.Front(); playerElement != nil; playerElement = playerElement.Next() {
		playerToExecuteActionFor := playerElement.Value.(*player.Player)

		err := action(playerToExecuteActionFor)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PlayerManager) CurrentPlayer() *player.Player {
	currentPlayer := p.currentPlayerElement.Value.(*player.Player)
	return currentPlayer
}

func (p *PlayerManager) ResetCurrentPlayer() {
	p.currentPlayerElement = p.players.Front()
}

func (p *PlayerManager) HasMorePlayers() bool {
	return p.currentPlayerElement != nil && p.currentPlayerElement.Next() != nil
}

func (p *PlayerManager) NextPlayer() {
	if p.currentPlayerElement != nil {
		p.currentPlayerElement = p.currentPlayerElement.Next()
	}
}

func (p *PlayerManager) NextStartPlayerAndReset() {
	p.players.MoveToBack(p.players.Front())
	p.currentPlayerElement = p.players.Front()
}
