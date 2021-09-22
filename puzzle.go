package knight

import (
	"fmt"
	"net/url"
)

func (a *Api) GetDailyPuzzle() (*DailyPuzzle, error) {
	puzzle := new(DailyPuzzle)
	err := a.get("api/puzzle/daily", AcceptNdjson, nil, puzzle)
	if err != nil {
		return nil, err
	}
	return puzzle, nil
}

func (a *Api) GetPuzzleActivity() (*PuzzleActivity, error) {
	puzzle := new(PuzzleActivity)

	u := make(url.Values)
	u.Add("max", "1") // TODO need to change (find the best way), because i dont get objects array in response

	err := a.get("api/puzzle/activity", AcceptNdjson, u, puzzle)
	if err != nil {
		return nil, err
	}
	return puzzle, nil
}

func (a *Api) GetPuzzleDashboard(days int) (*PuzzleDashboard, error) {
	puzzle := new(PuzzleDashboard)

	err := a.get(fmt.Sprintf("api/puzzle/dashboard/%v", days), AcceptJson, nil, puzzle)
	if err != nil {
		return nil, err
	}
	return puzzle, nil
}

func (a *Api) GetStormDashboard(username string) (*StormDashboard, error) {
	puzzle := new(StormDashboard)

	err := a.get(fmt.Sprintf("api/storm/dashboard/%v", username), AcceptJson, nil, puzzle)
	if err != nil {
		return nil, err
	}
	return puzzle, nil
}
