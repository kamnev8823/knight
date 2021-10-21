package knight

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const Host = "lichess.org"

const (
	AcceptJson    = "application/json"
	AcceptVndJson = "application/vnd.lichess.v3+json"
	AcceptNdjson  = "application/x-ndjson"
	AcceptPgn     = "application/x-chess-pgn"
	AcceptText    = "text/plain"
)

type Api struct {
	token string
}

func NewApi(token string) *Api {
	return &Api{token}
}

func (a *Api) get(endpoint, accept string, query url.Values, result interface{}) error {
	return a.call(http.MethodGet, endpoint, accept, query, nil, result)
}

// TODO find the best for getting pgn or plaintext value
func (a *Api) getPlain(endpoint, accept string, query url.Values) ([]byte, error) {
	res, err := a.callResponse(http.MethodGet, endpoint, accept, query, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func (a *Api) post(endpoint, accept string, query url.Values, body io.Reader, result interface{}) error {
	return a.call(http.MethodPost, endpoint, accept, query, body, result)
}

func (a *Api) delete(endpoint, accept string, query url.Values, result interface{}) error {
	return a.call(http.MethodDelete, endpoint, accept, query, nil, result)
}

func (a *Api) getEvent(endpoint string, query url.Values, result interface{}, si streamInterface) error {
	res, err := a.callResponse(http.MethodGet, endpoint, AcceptNdjson, query, nil)
	if err != nil {
		return err
	}

	go writeEventData(si, res, result)

	return nil
}

func (a *Api) postEvent(endpoint string, query url.Values, body io.Reader, result interface{}, si streamInterface) error {
	res, err := a.callResponse(http.MethodPost, endpoint, AcceptNdjson, query, body)
	if err != nil {
		return err
	}

	go writeEventData(si, res, result)

	return nil
}

//formRequest generate a request
func (a *Api) formRequest(method, endpoint string, acceptType string, query url.Values, body io.Reader) (*http.Request, error) {
	u := url.URL{
		Host:     Host,
		Path:     endpoint,
		RawQuery: query.Encode(),
		Scheme:   "https",
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token))
	req.Header.Add("Accept", acceptType)

	return req, nil
}

//callGetResponse send a request, return response and don't close connection
func (a *Api) callResponse(method, endpoint, acceptType string, query url.Values, body io.Reader) (*http.Response, error) {
	req, err := a.formRequest(method, endpoint, acceptType, query, body)
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
func (a *Api) call(method, endpoint, acceptType string, query url.Values, body io.Reader, result interface{}) error {
	req, err := a.formRequest(method, endpoint, acceptType, query, body)
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

func writeEventData(si streamInterface, response *http.Response, result interface{}) {
	defer response.Body.Close()
	defer si.close()

	for {
		if err := json.NewDecoder(response.Body).Decode(result); err != nil {
			break
		} else if ok := si.write(result); !ok {
			break
		}
	}

	return
}
