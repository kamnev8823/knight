package knight

func (a *Api) GetCurrentTVGames() (*TV, error) {
	tv := new(TV)
	err := a.get("api/tv/channels", CONTENT_TYPE_JSON, nil, tv)
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
	err := a.getEvent("api/tv/feed", CONTENT_TYPE_JSON, nil, v, tvc)
	if err != nil {
		return nil, err
	}
	return tvc, nil
}
