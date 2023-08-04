package main

import (
	"strings"
)

func newGame() GameState {
	return GameState{
		ListOfKills:    nil,
		ListOfPlayers:  nil,
		ListOfGunKills: nil,
		GunKills:       make(map[string]int),
		PlayerKills:    make(map[string]int)}
}

func getplayer2name(log []string, index int) string {
	// Handling names with spaces
	name := log[index]
	for log[index+1] != "by" {
		index++
		name = name + " " + log[index]
	}
	return name
}

func getplayer1name(log []string, index int) string {
	// Handling names with spaces
	name := log[index]
	for !strings.Contains(log[index-1], ":") {
		index--
		name = log[index] + " " + name
	}
	return name
}

func parse(log []string) Report {
	report := Report{}
	game := Game{}
	gameState := newGame()

	for i := 0; i < len(log); i++ {
		switch log[i] {
		case Token(GAME_SHUTDOWN).String():
			game.GameId = game.GameId + 1
			game.Players = gameState.ListOfPlayers
			for name, kills := range gameState.PlayerKills {
				gameState.ListOfKills = append(gameState.ListOfKills, PlayerNameKills{Name: name, Kills: kills})
			}
			for name, kills := range gameState.GunKills {
				gameState.ListOfGunKills = append(gameState.ListOfGunKills, GunTypeKills{Name: name, Kills: kills})
			}
			game.Kills = gameState.ListOfKills
			game.Kill_by_means = gameState.ListOfGunKills
			report.Report = append(report.Report, game)

		case Token(GAME_STARTED).String():
			gameState = newGame()
		case Token(CAUSE_BY).String():
			gunName := log[i+1]
			gameState.GunKills[gunName] = gameState.GunKills[gunName] + 1
		case Token(KILL).String():
			player_1 := getplayer1name(log, (i - 1))
			player_2 := getplayer2name(log, (i + 1))

			if !isPlayerInList(gameState.ListOfPlayers, player_1) && player_1 != "<world>" {
				gameState.ListOfPlayers = append(gameState.ListOfPlayers, player_1)
			}
			if !isPlayerInList(gameState.ListOfPlayers, player_2) {
				gameState.ListOfPlayers = append(gameState.ListOfPlayers, player_2)
			}
			if player_1 == "<world>" {
				gameState.PlayerKills[player_2] = gameState.PlayerKills[player_2] - 1
				game.Total_kills = game.Total_kills - 1
			} else {
				gameState.PlayerKills[player_1] = gameState.PlayerKills[player_1] + 1
				game.Total_kills = game.Total_kills + 1
			}
		}
	}
	return report
}
