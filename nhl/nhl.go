package nhl

import (
	"fmt"
	"time"

	mysportsfeeds "github.com/delaneyj/mysportsfeeds-go"
)

func seasonInfoToName(seasonStart int, isRegular bool) string {
	if seasonStart < 100 {
		seasonStart += 2000
	}
	postFix := "regular"
	if !isRegular {
		postFix = "playoff"
	}
	return fmt.Sprintf("%d-%d-%s", seasonStart, seasonStart+1, postFix)
}

//NHL x
type NHL struct {
	seasonName string
	c          mysportsfeeds.Client
}

// NewNHL x
func NewNHL(c mysportsfeeds.Client, seasonStart int, isPlayoffs bool) *NHL {
	nhl := NHL{
		seasonInfoToName(seasonStart, isPlayoffs),
		c,
	}
	return &nhl
}

type cumulativePlayerStatsResponse struct {
	CumulativePlayerStats *CumulativePlayerStats `json:"cumulativeplayerstats"`
}

//Player x
type Player struct {
	ID           int    `json:"ID"`
	LastName     string `json:"LastName"`
	FirstName    string `json:"FirstName"`
	JerseyNumber int    `json:"JerseyNumber"`
	Position     string `json:"Position"`
	Height       string `json:"Height"`
	Weight       string `json:"Weight"`
	BirthDate    string `json:"BirthDate"`
	Age          int    `json:"Age"`
	BirthCity    string `json:"BirthCity"`
	BirthCountry string `json:"BirthCountry"`
	IsRookie     bool   `json:"IsRookie"`
}

//Team x
type Team struct {
	ID           int    `json:"ID"`
	City         string `json:"City"`
	Name         string `json:"Name"`
	Abbreviation string `json:"Abbreviation"`
}

//Stats x
type Stats struct {
}

//PlayerStatsEntry x
type PlayerStatsEntry struct {
	Player *Player `json:"player"`
	Team   *Team   `json:"team"`
	Stats  *Stats  `json:"stats"`
}

//CumulativePlayerStats x
type CumulativePlayerStats struct {
	LastUpdatedOn time.Time           `json:"lastUpdatedOn"`
	PlayersStats  []*PlayerStatsEntry `json:"playerstatsentry"`
}

//CumulativePlayerStats x
func (nhl *NHL) CumulativePlayerStats() (*CumulativePlayerStats, error) {
	url := fmt.Sprintf("nhl/%s/cumulative_player_stats.%s", nhl.seasonName, mysportsfeeds.Format)

	var cps CumulativePlayerStats
	_, err := nhl.c.Request(url, &cps)
	if err != nil {
		return nil, err
	}

	return &cps, nil
}
