package knight

import (
	"fmt"
	"net/url"
	"strconv"
)

type exportGame struct {
	Json *Game
	Pgn  []byte
}

func (a *Api) ExportGameJson(gameId string, moves, pgnInJson, tags, clocks, opening bool) (*Game, error) {
	route := fmt.Sprintf("game/export/%v", gameId)
	r, err := a.exportJsonOrPgn(route, moves, pgnInJson, tags, clocks, opening, false)
	if err != nil {
		return nil, err
	}
	return r.Json, nil
}

func (a *Api) ExportGamePgn(gameId string, moves, pgnInJson, tags, clocks, opening bool) ([]byte, error) {
	route := fmt.Sprintf("game/export/%v", gameId)
	r, err := a.exportJsonOrPgn(route, moves, pgnInJson, tags, clocks, opening, true)
	if err != nil {
		return nil, err
	}
	return r.Pgn, nil
}

func (a *Api) ExportOngoingGameJson(username string, moves, pgnInJson, tags, clocks, opening bool) (*Game, error) {
	route := fmt.Sprintf("api/user/%v/current-game", username)
	r, err := a.exportJsonOrPgn(route, moves, pgnInJson, tags, clocks, opening, false)
	if err != nil {
		return nil, err
	}
	return r.Json, nil
}

func (a *Api) ExportOngoingGamePgn(username string, moves, pgnInJson, tags, clocks, opening bool) ([]byte, error) {
	route := fmt.Sprintf("api/user/%v/current-game", username)
	r, err := a.exportJsonOrPgn(route, moves, pgnInJson, tags, clocks, opening, true)
	if err != nil {
		return nil, err
	}
	return r.Pgn, nil
}

func (a *Api) exportJsonOrPgn(route string, moves, pgnInJson, tags, clocks, opening, isPgn bool) (*exportGame, error) {
	game := &exportGame{
		new(Game),
		[]byte{},
	}

	u := make(url.Values)
	u.Add("tags", strconv.FormatBool(tags))
	u.Add("moves", strconv.FormatBool(moves))
	u.Add("clocks", strconv.FormatBool(clocks))
	u.Add("opening", strconv.FormatBool(opening))
	u.Add("pgnInJson", strconv.FormatBool(pgnInJson))

	if isPgn {
		l, err := a.getPlain(route, AcceptPgn, nil)
		if err != nil {
			return nil, err
		}
		game.Pgn = l

	} else if err := a.get(route, AcceptJson, nil, game.Json); err != nil {
		return nil, err
	}

	return game, nil
}
