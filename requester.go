package squarecloud

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/End313234/squarecloud-wrapper/constants"
)

type requester struct {
	client        http.Client
	Authorization string
}

func newRequester(authorization string) *requester {
	return &requester{
		client:        http.Client{},
		Authorization: authorization,
	}
}

func (r *requester) Get(route string, v any) error {
	request, _ := http.NewRequest(http.MethodGet, constants.BASE_URL+route, nil)
	request.Header["Authorization"] = []string{r.Authorization}

	response, err := r.client.Do(request)
	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		var err Error
		json.Unmarshal(body, &err)

		return errors.New(err.Code)
	}

	json.Unmarshal(body, v)
	return nil
}
