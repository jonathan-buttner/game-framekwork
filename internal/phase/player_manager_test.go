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
	actionFunc := func(state phase.ExecutionState) error {
		playerNames[state.CurrentPlayer.Name] = struct{}{}
		return nil
	}

	manager.ExecuteForPlayers(actionFunc)

	assert.Equal(t, len(playerNames), 2)
}

func TestExecutesForAllPlayersPrevPlayerIsNilForFirstPlayer(t *testing.T) {
	player1 := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player2 := player.NewPlayer("player2", rules.NewDefaultGameRules())
	manager := phase.NewPlayerManager([]*player.Player{player1, player2})

	pos := 0
	actionFunc := func(state phase.ExecutionState) error {
		defer func() { pos++ }()

		assert.Equal(t, state.Position, pos)
		if pos == 0 {
			assert.Equal(t, state.PrevPlayer.Name, "player2")
			assert.Equal(t, state.CurrentPlayer.Name, "player1")
			assert.Equal(t, state.NextPlayer.Name, "player2")

		} else if pos == 1 {
			assert.Equal(t, state.CurrentPlayer.Name, "player2")
			assert.Equal(t, state.PrevPlayer.Name, "player1")
			assert.Equal(t, state.NextPlayer.Name, "player1")

		}
		return nil
	}

	manager.ExecuteForPlayers(actionFunc)
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

	manager.RotateStartPlayer()
	assert.Equal(t, manager.CurrentPlayer().Name, "player2")
}

func TestNextStartPlayerIsPlayer2AfterReset(t *testing.T) {
	player1 := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player2 := player.NewPlayer("player2", rules.NewDefaultGameRules())
	manager := phase.NewPlayerManager([]*player.Player{player1, player2})

	manager.RotateStartPlayer()
	manager.NextPlayer()
	manager.ResetCurrentPlayer()
	assert.Equal(t, manager.CurrentPlayer().Name, "player2")
}
