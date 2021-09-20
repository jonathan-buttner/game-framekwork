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

type ExecutionState struct {
	CurrentPlayer *player.Player
	NextPlayer    *player.Player
	PrevPlayer    *player.Player
	Position      int
}

type PlayersAction func(state ExecutionState) error

func (p *PlayerManager) ExecuteForPlayers(action PlayersAction) error {
	pos := 0

	for playerElement := p.players.Front(); playerElement != nil; playerElement = playerElement.Next() {
		playerToExecuteActionFor := playerElement.Value.(*player.Player)

		nextPlayer := p.getNextPlayer(playerElement)
		prevPlayer := p.getPrevPlayer(playerElement)

		err := action(ExecutionState{
			CurrentPlayer: playerToExecuteActionFor,
			NextPlayer:    nextPlayer,
			PrevPlayer:    prevPlayer,
			Position:      pos,
		})
		if err != nil {
			return err
		}

		pos++
	}

	return nil
}

func (p *PlayerManager) getNextPlayer(playerElement *list.Element) *player.Player {
	nextPlayerElement := playerElement.Next()
	if nextPlayerElement == nil {
		// this means we're at the back of the list, so wrap around and get the first player
		return p.players.Front().Value.(*player.Player)
	}

	return nextPlayerElement.Value.(*player.Player)
}

func (p *PlayerManager) getPrevPlayer(playerElement *list.Element) *player.Player {
	prevPlayerElement := playerElement.Prev()
	if prevPlayerElement == nil {
		// this means we're at the front of the list, so wrap around and get the laster player
		return p.players.Back().Value.(*player.Player)
	}

	return prevPlayerElement.Value.(*player.Player)
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

func (p *PlayerManager) RotateStartPlayer() {
	p.players.MoveToBack(p.players.Front())
	p.currentPlayerElement = p.players.Front()
}
