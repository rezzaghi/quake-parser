package main

type PlayerNameKills struct {
	Name  string
	Kills int
}

type GunTypeKills struct {
	Name  string
	Kills int
}

type Game struct {
	GameId        int
	Total_kills   int
	Players       []string
	Kills         [](PlayerNameKills)
	Kill_by_means [](GunTypeKills)
}

type Report struct {
	Report []Game
}

type GameState struct {
	ListOfPlayers  []string
	PlayerKills    map[string]int
	GunKills       map[string]int
	ListOfKills    []PlayerNameKills
	ListOfGunKills []GunTypeKills
}
