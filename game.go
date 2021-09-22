package knight

import (
	"fmt"
	"net/url"
	"strconv"
)

type exportGame struct {
	json *Game
	pgn  []byte
}

func (a *Api) ExportGame(gameId string, moves, pgnInJson, tags, clocks, opening, isPgn bool) (*Game, error) {
	route := fmt.Sprintf("game/export/%v", gameId)
	eg, err := a.exportJsonOrPgn(route, moves, pgnInJson, tags, clocks, opening, isPgn)
	if err != nil {
		return nil, err
	}
	return eg.json, nil
}

func (a *Api) ExportOngoingGame(username string, moves, pgnInJson, tags, clocks, opening, isPgn bool) ([]byte, error) {
	route := fmt.Sprintf("api/user/%v/current-game", username)
	eg, err := a.exportJsonOrPgn(route, moves, pgnInJson, tags, clocks, opening, isPgn)
	if err != nil {
		return nil, err
	}
	return eg.pgn, nil
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
		game.pgn = l

	} else if err := a.get(route, AcceptJson, nil, game.json); err != nil {
		return nil, err
	}

	return game, nil
}
