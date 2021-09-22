package knight

import (
	"fmt"
	"net/url"
	"strconv"
)

func (a *Api) GetCurrentTVGames() (*TV, error) {
	tv := new(TV)
	err := a.get("api/tv/channels", AcceptJson, nil, tv)
	if err != nil {
		return nil, err
	}
	return tv, nil
}

func (a *Api) StreamCurrentTVGame() (*TVChan, error) {
	tvc := &TVChan{
		channel: make(chan *TVStream),
		closed:  false,
	}
	v := new(TVStream)
	err := a.getEvent("api/tv/feed", nil, v, tvc)
	if err != nil {
		return nil, err
	}
	return tvc, nil
}

func (a *Api) GetBestTVOngoingGames(channel string, moves, pgnInJson, tags, clocks, opening bool) (*Game, error) {
	best := new(Game)

	u := make(url.Values)
	u.Add("nb", "1") // TODO need to change (find the best way), because i dont get objects array in response
	u.Add("tags", strconv.FormatBool(tags))
	u.Add("moves", strconv.FormatBool(moves))
	u.Add("clocks", strconv.FormatBool(clocks))
	u.Add("opening", strconv.FormatBool(opening))
	u.Add("pgnInJson", strconv.FormatBool(pgnInJson))

	err := a.get(fmt.Sprintf("api/tv/%v", channel), AcceptNdjson, u, best)
	if err != nil {
		return nil, err
	}
	return best, err
}

func (a *Api) GetBestTVOngoingGamesPGN(channel string, nb int, moves, pgnInJson, tags, clocks, opening bool) ([]byte, error) {
	u := make(url.Values)
	u.Add("nb", strconv.Itoa(nb))
	u.Add("tags", strconv.FormatBool(tags))
	u.Add("moves", strconv.FormatBool(moves))
	u.Add("clocks", strconv.FormatBool(clocks))
	u.Add("opening", strconv.FormatBool(opening))
	u.Add("pgnInJson", strconv.FormatBool(pgnInJson))

	return a.getPlain(fmt.Sprintf("api/tv/%v", channel), AcceptPgn, u)
}
