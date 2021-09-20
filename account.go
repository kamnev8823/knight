package knight

import (
	"fmt"
	"net/url"
)

func (a *Api) GetProfile() (*Account, error) {
	account := new(Account)
	err := a.get("/api/account", CONTENT_TYPE_JSON, nil, account)
	if err != nil {
		return nil, err
	}
	return account, err
}

func (a *Api) GetEmail() (*AccountEmail, error) {
	email := new(AccountEmail)
	err := a.get("/api/account/email", CONTENT_TYPE_JSON, nil, email)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func (a *Api) GetPreference() (*Preferences, error) {
	preferences := new(Preferences)
	err := a.get("/api/account/preferences", CONTENT_TYPE_JSON, nil, preferences)
	if err != nil {
		return nil, err
	}
	return preferences, nil
}

func (a *Api) GetKidMode() (*KidMode, error) {
	kid := new(KidMode)
	err := a.get("/api/account/kid", CONTENT_TYPE_JSON, nil, kid)
	if err != nil {
		return nil, err
	}
	return kid, nil
}

func (a *Api) SetKidMode(v bool) (bool, error) {
	u := make(url.Values)
	u.Add("v", fmt.Sprintf("%v", v))

	isOk := new(isOk)
	err := a.post("/api/account/kid", CONTENT_TYPE_JSON, u, nil, isOk)
	if err != nil {
		return false, err
	}

	return isOk.Ok, nil
}
