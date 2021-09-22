package knight

import "fmt"

func (a *Api) GetFollowing(username string) (*Account, error) {
	foll := new(Account)
	err := a.get(fmt.Sprintf("api/user/%v/following", username), AcceptNdjson, nil, foll)
	if err != nil {
		return nil, err
	}
	return foll, nil
}

func (a *Api) GetFollowers(username string) (*Account, error) {
	foll := new(Account)
	err := a.get(fmt.Sprintf("api/user/%v/followers", username), AcceptJson, nil, foll)
	if err != nil {
		return nil, err
	}
	return foll, nil
}

func (a *Api) FollowPlayer(username string) (bool, error) {
	ok := new(isOk)
	err := a.post(fmt.Sprintf("api/rel/follow/%v", username), AcceptJson, nil, nil, ok)
	if err != nil {
		return false, err
	}
	return ok.Ok, err
}

func (a *Api) UnfollowPlayer(username string) (bool, error) {
	ok := new(isOk)
	err := a.post(fmt.Sprintf("api/rel/unfollow/%v", username), AcceptJson, nil, nil, ok)
	if err != nil {
		return false, err
	}
	return ok.Ok, err
}
