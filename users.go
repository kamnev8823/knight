package knight

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
)

func (a *Api) GetUsersStatus(ids ...string) ([]*UserStatus, error) {
	var usersStatus []*UserStatus

	u := make(url.Values)
	u.Add("ids", strings.Join(ids, ","))

	err := a.get("api/users/status", u, &usersStatus)
	if err != nil {
		return nil, err
	}
	return usersStatus, nil
}

func (a *Api) GetUser(username string) (*Account, error) {
	account := new(Account)

	err := a.get(fmt.Sprintf("api/user/%v", username), nil, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Api) GetUserHistory(username string) ([]*History, error) {
	var history []*History

	err := a.get(fmt.Sprintf("api/user/%v/rating-history", username), nil, &history)
	if err != nil {
		return nil, err
	}

	return history, nil
}

//GetPerformance Read performance statistics of a user, for a single performance.
//
//perf - Enum:
// 	"ultraBullet" "bullet" "blitz" "rapid"
//	"classical" "correspondence" "chess960"
//	"crazyhouse" "antichess" "atomic" "horde"
//	"kingOfTheHill" "racingKings" "threeCheck"
func (a *Api) GetPerformance(username string, perf string) (*Performance, error) {
	performance := new(Performance)

	err := a.get(fmt.Sprintf("api/user/%v/perf/%v", username, perf), nil, performance)
	if err != nil {
		return nil, err
	}

	return performance, nil
}

func (a *Api) GetUserActivity(username string) ([]*Activity, error) {
	var activity []*Activity

	err := a.get(fmt.Sprintf("api/user/%v/activity", username), nil, &activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (a *Api) GetUsersById(ids []string) ([]*User, error) {
	var users []*User

	body := bytes.NewReader([]byte(strings.Join(ids, ",")))
	err := a.post("api/users", nil, body, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *Api) GetMembersTeam(teamId string) (*Account, error) {
	member := new(Account)
	err := a.get(fmt.Sprintf("api/team/%v/users", teamId), nil, member)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (a *Api) GetLiveStreamer() ([]*Streamer, error) {
	var streamers []*Streamer

	err := a.get("streamer/live", nil, &streamers)
	if err != nil {
		return nil, err
	}
	return streamers, nil
}

func (a *Api) GetCrosstable(user1 string, user2 string, matchup bool) (*Crosstable, error) {
	crosstable := new(Crosstable)

	u := make(url.Values)
	u.Add("mathcup", fmt.Sprintf("%v", matchup))

	err := a.get(fmt.Sprintf("api/crosstable/%v/%v", user1, user2), u, crosstable)
	if err != nil {
		return nil, err
	}
	return crosstable, err
}
