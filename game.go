package knight

import (
	"fmt"
	"net/url"
	"strconv"
)

type ExportGame struct {
	Json *Game
	Pgn  []byte
}

func (a *Api) ExportGame(gameId string, moves, pgnInJson, tags, clocks, opening, isPgn bool) (*ExportGame, error) {
	route := fmt.Sprintf("game/export/%v", gameId)
	return a.exportJsonOrPgn(route, moves, pgnInJson, tags, clocks, opening, isPgn)
}

func (a *Api) ExportOngoingGame(username string, moves, pgnInJson, tags, clocks, opening, isPgn bool) (*ExportGame, error) {
	route := fmt.Sprintf("api/user/%v/current-game", username)
	return a.exportJsonOrPgn(route, moves, pgnInJson, tags, clocks, opening, isPgn)
}

func (a *Api) exportJsonOrPgn(route string, moves, pgnInJson, tags, clocks, opening, isPgn bool) (*ExportGame, error) {
	game := &ExportGame{
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
