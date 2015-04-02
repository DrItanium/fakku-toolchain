package fakku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ResponseOk                         = 200
	ResponseNotFound                   = 404
	ResponseUnavailableForLegalReasons = 451 // got DMCA'd son
	ResponseServiceUnavailable         = 503 // Information unavailable
	apiHeader                          = "https://api.fakku.net/"
)

type apiFunction interface {
	Construct() string
}

type supportsPagination struct {
	Page uint
}

type ErrorStatus struct {
	ErrorCode    int
	ErrorMessage string `json:"error"`
	KnownError   bool
}

func (e ErrorStatus) Error() string {
	return fmt.Sprintf("Error %d: %s", e.ErrorCode, e.ErrorMessage)
}

type UnknownEntry struct {
	Message string
}

func (e UnknownEntry) Error() string {
	return e.Message
}

func apiCall(url apiFunction, c interface{}) error {
	resp, err := http.Get(url.Construct())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case ResponseOk:
		err = json.Unmarshal(body, &c)
		if err != nil {
			return err
		}
		return nil
	case ResponseNotFound, ResponseUnavailableForLegalReasons, ResponseServiceUnavailable:
		// right now just harvest the error code
		var ec ErrorStatus
		err = json.Unmarshal(body, &ec)
		ec.KnownError = true
		ec.ErrorCode = resp.StatusCode
		if err != nil {
			return err
		} else {
			return &ec
		}
	default:
		return &ErrorStatus{ErrorCode: resp.StatusCode, ErrorMessage: resp.Status, KnownError: false}
	}
}
func paginateString(s string, page uint) string {
	// If page is zero then it is meaningless so just return the string
	if page == 0 {
		return s
	} else {
		return fmt.Sprintf("%s/page/%d", s, page)
	}
}
