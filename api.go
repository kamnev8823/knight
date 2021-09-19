package knight

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const HOST = "lichess.org"

type Api struct {
	token string
}

func NewApi(token string) *Api {
	return &Api{token}
}

func (a *Api) get(endpoint string, query url.Values, result interface{}) error {
	return a.call(http.MethodGet, endpoint, query, nil, result)
}

func (a *Api) post(endpoint string, query url.Values, body io.Reader, result interface{}) error {
	return a.call(http.MethodPost, endpoint, query, body, result)
}

func (a *Api) delete(endpoint string, query url.Values, result interface{}) error {
	return a.call(http.MethodDelete, endpoint, query, nil, result)
}

func (a *Api) getEvent(endpoint string, query url.Values, result interface{}, sei streamEventInterface) error {
	res, err := a.callResponse(http.MethodGet, endpoint, query, nil)
	if err != nil {
		return err
	}

	go writeEventData(sei, res, result)

	return nil
}

func (a *Api) postEvent(endpoint string, query url.Values, body io.Reader, result interface{}, sei streamEventInterface) error {
	res, err := a.callResponse(http.MethodPost, endpoint, query, body)
	if err != nil {
		return err
	}

	go writeEventData(sei, res, result)

	return nil
}

//formRequest generate a request
func (a *Api) formRequest(method, endpoint string, query url.Values, body io.Reader) (*http.Request, error) {
	u := url.URL{
		Host:     HOST,
		Path:     endpoint,
		RawQuery: query.Encode(),
		Scheme:   "https",
	}
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token))
	return req, nil
}

//callGetResponse send a request, return response and don't close connection
func (a *Api) callResponse(method, endpoint string, query url.Values, body io.Reader) (*http.Response, error) {
	req, err := a.formRequest(method, endpoint, query, body)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//callDontClose send a request and write the result
func (a *Api) call(method, endpoint string, query url.Values, body io.Reader, result interface{}) error {
	req, err := a.formRequest(method, endpoint, query, body)
	if err != nil {
		return err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// TODO change
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		lichessErr := &Error{}
		if err := json.NewDecoder(res.Body).Decode(lichessErr); err != nil {
			return err
		}
		return errors.New(lichessErr.Message)
	}

	if result != nil {
		if err := json.NewDecoder(res.Body).Decode(result); err != nil {
			return err
		}
	}
	return nil
}

func writeEventData(sei streamEventInterface, response *http.Response, result interface{}) {
	defer response.Body.Close()

	for {
		if sei.isClosed() {
			break
		} else if err := json.NewDecoder(response.Body).Decode(result); err == nil {
			sei.write(result)
		} else {
			sei.Close()
			break
		}
	}

	return
}
