package phase_test

import (
	"testing"

	"github.com/jonathan-buttner/game-framework/internal/phase"
	"github.com/jonathan-buttner/game-framework/internal/player"
	"github.com/jonathan-buttner/game-framework/internal/rules"
	"github.com/stretchr/testify/assert"
)

func TestExecutesForAllPlayers(t *testing.T) {
	player1 := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player2 := player.NewPlayer("player2", rules.NewDefaultGameRules())
	manager := phase.NewPlayerManager([]*player.Player{player1, player2})

	playerNames := make(map[string]struct{})
	actionFunc := func(execPlayer *player.Player) error {
		playerNames[execPlayer.Name] = struct{}{}
		return nil
	}

	manager.ExecuteForPlayers(actionFunc)

	assert.Equal(t, len(playerNames), 2)
}

func TestInitialStartPlayerIsFirstInArray(t *testing.T) {
	player1 := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player2 := player.NewPlayer("player2", rules.NewDefaultGameRules())
	manager := phase.NewPlayerManager([]*player.Player{player1, player2})

	assert.Equal(t, manager.CurrentPlayer().Name, "player1")
}

func TestHasMorePlayersIsTrueInitiallyWithTwoPlayers(t *testing.T) {
	player1 := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player2 := player.NewPlayer("player2", rules.NewDefaultGameRules())
	manager := phase.NewPlayerManager([]*player.Player{player1, player2})

	assert.True(t, manager.HasMorePlayers())
}

func TestHasMorePlayersIsFalseAfterNextPlayer(t *testing.T) {
	player1 := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player2 := player.NewPlayer("player2", rules.NewDefaultGameRules())
	manager := phase.NewPlayerManager([]*player.Player{player1, player2})

	manager.NextPlayer()
	assert.False(t, manager.HasMorePlayers())
}

func TestResetsToPlayer1(t *testing.T) {
	player1 := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player2 := player.NewPlayer("player2", rules.NewDefaultGameRules())
	manager := phase.NewPlayerManager([]*player.Player{player1, player2})

	manager.NextPlayer()
	manager.ResetCurrentPlayer()
	assert.Equal(t, manager.CurrentPlayer().Name, "player1")
}

func TestNextStartPlayerIsPlayer2(t *testing.T) {
	player1 := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player2 := player.NewPlayer("player2", rules.NewDefaultGameRules())
	manager := phase.NewPlayerManager([]*player.Player{player1, player2})

	manager.NextStartPlayerAndReset()
	assert.Equal(t, manager.CurrentPlayer().Name, "player2")
}

func TestNextStartPlayerIsPlayer2AfterReset(t *testing.T) {
	player1 := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player2 := player.NewPlayer("player2", rules.NewDefaultGameRules())
	manager := phase.NewPlayerManager([]*player.Player{player1, player2})

	manager.NextStartPlayerAndReset()
	manager.NextPlayer()
	manager.ResetCurrentPlayer()
	assert.Equal(t, manager.CurrentPlayer().Name, "player2")
}
