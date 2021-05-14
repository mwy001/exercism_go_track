package tournament

import (
	"io"
	"bufio"
	"strings"
	"errors"
	"sort"
	"fmt"
)


type Team struct {
	Name string
	MatchesPlayed int
	MatchesWon int
	MatchesDrawn int
	MatchesLost int
	Points int
}

func (team *Team) String() string {
	return fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d", team.Name, team.MatchesPlayed, team.MatchesWon, team.MatchesDrawn, team.MatchesLost, team.Points)
}

func win(teams map[string]*Team, teamName string) {
	team, ok := teams[teamName]
	if !ok {
		team = &Team{}
		team.Name = teamName
		teams[teamName] = team
	}

	team.MatchesPlayed += 1
	team.MatchesWon += 1
	team.Points += 3
}

func lost(teams map[string]*Team, teamName string) {
	team, ok := teams[teamName]
	if !ok {
		team = &Team{}
		team.Name = teamName
		teams[teamName] = team
	}

	team.MatchesPlayed += 1
	team.MatchesLost += 1
}

func tie(teams map[string]*Team, teamName string) {
	team, ok := teams[teamName]
	if !ok {
		team = &Team{}
		team.Name = teamName
		teams[teamName] = team
	}

	team.MatchesPlayed += 1
	team.MatchesDrawn += 1
	team.Points += 1
}

func Tally(input io.Reader, output io.Writer) error {
	var teams map[string]*Team = map[string]*Team{}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || line[0] == '#' {
			continue
		}

		parts := strings.Split(line, ";")

		if len(parts) != 3 || (parts[2] != "win" && parts[2] != "loss" && parts[2] != "draw") {
			return errors.New("malformatted input line")
		}

		firstTeam := parts[0]
		secondTeam := parts[1]
		result := parts[2]

		switch result {
		case "win":
			win(teams, firstTeam)
			lost(teams, secondTeam)
		case "loss":
			win(teams, secondTeam)
			lost(teams, firstTeam)
		case "draw":
			tie(teams, firstTeam)
			tie(teams, secondTeam)	
		default:
		}
	}
	var teamsSlice []*Team
	for _, v := range teams {
		teamsSlice = append(teamsSlice, v)
	}

	sort.Slice(teamsSlice, func(i, j int) bool {
		if teamsSlice[i].Points != teamsSlice[j].Points {
			return teamsSlice[i].Points > teamsSlice[j].Points
		} else {
			return teamsSlice[i].Name < teamsSlice[j].Name
		}
	})

	output.Write([]byte("Team                           | MP |  W |  D |  L |  P"))
	output.Write([]byte("\n"))
	
	for _, v := range teamsSlice {
		output.Write([]byte(v.String()))
		output.Write([]byte("\n"))
	}

	return nil
}
