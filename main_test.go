package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	stringTest := []string{"InitGame:", "Roberto", "killed", "Carlos", "by", "GUN", "ShutdownGame:"}

	got := parse(stringTest)

	playersKills := []PlayerNameKills{
		{Name: "Roberto", Kills: 1},
	}
	gunkills := []GunTypeKills{{Name: "GUN", Kills: 1}}
	list := []string{"Roberto", "Carlos"}
	report := Report{Report: []Game{{GameId: 1, Total_kills: 1, Kills: playersKills, Players: list, Kill_by_means: gunkills}}}

	if got.Report[0].GameId != report.Report[0].GameId {
		t.Errorf("got %q, wanted %q", got.Report[0].GameId, report.Report[0].GameId)
	}

	if got.Report[0].Total_kills != report.Report[0].Total_kills {
		t.Errorf("got %q, wanted %q", got.Report[0].Total_kills, report.Report[0].Total_kills)
	}

	if got.Report[0].Players[0] != report.Report[0].Players[0] {
		t.Errorf("got %q, wanted %q", got.Report[0].Players[0], report.Report[0].Players[0])
	}

	if got.Report[0].Kills[0].Kills != report.Report[0].Kills[0].Kills {
		t.Errorf("got %q, wanted %q", got.Report[0].Kills[0].Kills, report.Report[0].Kills[0].Kills)
	}

	if got.Report[0].Kills[0].Name != report.Report[0].Kills[0].Name {
		t.Errorf("got %q, wanted %q", got.Report[0].Kills[0].Name, report.Report[0].Kills[0].Name)
	}

	if got.Report[0].Kill_by_means[0].Kills != report.Report[0].Kill_by_means[0].Kills {
		t.Errorf("got %q, wanted %q", got.Report[0].Kill_by_means[0].Kills, report.Report[0].Kill_by_means[0].Kills)
	}

	if got.Report[0].Kill_by_means[0].Name != report.Report[0].Kill_by_means[0].Name {
		t.Errorf("got %q, wanted %q", got.Report[0].Kill_by_means[0].Name, report.Report[0].Kill_by_means[0].Name)
	}

}
