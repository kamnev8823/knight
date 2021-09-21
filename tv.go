package knight

import (
	"fmt"
	"net/url"
)

func (a *Api) GetCurrentTVGames() (*TV, error) {
	tv := new(TV)
	err := a.get("api/tv/channels", nil, tv)
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

func (a *Api) GetBestTVOngoingGames(channel string, moves, pgnInJson, tags, clocks, opening bool) (*TVBest, error) {
	best := new(TVBest)

	u := make(url.Values)
	u.Add("nb", "1") // TODO need to change, because i dont get objects array in response
	u.Add("tags", fmt.Sprintf("%v", tags))
	u.Add("moves", fmt.Sprintf("%v", moves))
	u.Add("clocks", fmt.Sprintf("%v", clocks))
	u.Add("opening", fmt.Sprintf("%v", opening))
	u.Add("pgnInJson", fmt.Sprintf("%v", pgnInJson))

	err := a.get(fmt.Sprintf("api/tv/%v", channel), u, best)
	if err != nil {
		return nil, err
	}
	return best, err
}
