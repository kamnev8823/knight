package knight

import (
	"fmt"
	"net/url"
)

func (a *Api) GetDailyPuzzle() (*DailyPuzzle, error) {
	puzzle := new(DailyPuzzle)
	err := a.get("api/puzzle/daily", nil, puzzle)
	if err != nil {
		return nil, err
	}
	return puzzle, nil
}

func (a *Api) GetPuzzleActivity() (*PuzzleActivity, error) {
	puzzle := new(PuzzleActivity)

	u := make(url.Values)
	u.Add("max", "1") // TODO need to change, because i dont get objects array in response

	err := a.get("api/puzzle/activity", u, puzzle)
	if err != nil {
		return nil, err
	}
	return puzzle, nil
}

func (a *Api) GetPuzzleDashboard(days int) (*PuzzleDashboard, error) {
	puzzle := new(PuzzleDashboard)

	err := a.get(fmt.Sprintf("api/puzzle/dashboard/%v", days), nil, puzzle)
	if err != nil {
		return nil, err
	}
	return puzzle, nil
}

func (a *Api) GetStormDashboard(username string) (*StormDashboard, error) {
	puzzle := new(StormDashboard)

	err := a.get(fmt.Sprintf("api/storm/dashboard/%v", username), nil, puzzle)
	if err != nil {
		return nil, err
	}
	return puzzle, nil
}
