package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GamesResponse struct {
	Games []struct {
		Name    string `json:"name"`
		Checked int    `json:"checked"`
	} `json:"games"`
}

type Game struct {
	Name    string
	Checked bool
}

type BulletService struct {
	session string
}

func NewBulletService(session string) *BulletService {
	return &BulletService{
		session: session,
	}
}

func (bulletService *BulletService) getGamesRequest() (*GamesResponse, error) {
	request, err := http.NewRequest(http.MethodGet, "https://7bullet.de/api/winchallenge/list", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	request.AddCookie(&http.Cookie{Name: "connect.sid", Value: bulletService.session})
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer response.Body.Close()
	gamesResponse := new(GamesResponse)
	err = json.NewDecoder(response.Body).Decode(gamesResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}
	return gamesResponse, nil
}

func (bulletService *BulletService) GetGames() ([]Game, error) {
	rawGames, err := bulletService.getGamesRequest()
	if err != nil {
		return nil, fmt.Errorf("error getting games: %w", err)
	}
	games := []Game{}
	for _, rawGame := range rawGames.Games {
		if rawGame.Checked == 0 {
			games = append(games, Game{
				Name:    rawGame.Name,
				Checked: false,
			})
			continue
		}
		if rawGame.Checked == 1 {
			games = append(games, Game{
				Name:    rawGame.Name,
				Checked: true,
			})
			continue
		}
	}
	return games, nil
}
