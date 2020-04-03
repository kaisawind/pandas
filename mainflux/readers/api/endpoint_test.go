// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloustone/pandas/mainflux/readers"
	"github.com/cloustone/pandas/mainflux/readers/api"
	"github.com/cloustone/pandas/mainflux/readers/mocks"
	"github.com/cloustone/pandas/mainflux/transformers/senml"
	"github.com/cloustone/pandas/mainflux"
	"github.com/stretchr/testify/assert"
)

const (
	svcName       = "test-service"
	token         = "1"
	invalid       = "invalid"
	numOfMessages = 42
	chanID        = "1"
	valueFields   = 5
)

var (
	v       float64 = 5
	stringV         = "value"
	boolV           = true
	dataV           = "base64"
	sum     float64 = 42
)

func newService() readers.MessageRepository {
	messages := []senml.Message{}
	for i := 0; i < numOfMessages; i++ {
		msg := senml.Message{
			Channel:   chanID,
			Publisher: "1",
			Protocol:  "mqtt",
		}
		// Mix possible values as well as value sum.
		count := i % valueFields

		switch count {
		case 0:
			msg.Value = &v
		case 1:
			msg.BoolValue = &boolV
		case 2:
			msg.StringValue = &stringV
		case 3:
			msg.DataValue = &dataV
		case 4:
			msg.Sum = &sum
		}

		messages = append(messages, msg)
	}

	return mocks.NewMessageRepository(map[string][]senml.Message{
		chanID: messages,
	})
}

func newServer(repo readers.MessageRepository, tc mainflux.ThingsServiceClient) *httptest.Server {
	mux := api.MakeHandler(repo, tc, svcName)
	return httptest.NewServer(mux)
}

type testRequest struct {
	client *http.Client
	method string
	url    string
	token  string
}

func (tr testRequest) make() (*http.Response, error) {
	req, err := http.NewRequest(tr.method, tr.url, nil)
	if err != nil {
		return nil, err
	}
	if tr.token != "" {
		req.Header.Set("Authorization", tr.token)
	}

	return tr.client.Do(req)
}

func TestReadAll(t *testing.T) {
	svc := newService()
	tc := mocks.NewThingsService()
	ts := newServer(svc, tc)
	defer ts.Close()

	cases := map[string]struct {
		url    string
		token  string
		status int
	}{
		"read page with valid offset and limit": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=0&limit=10", ts.URL, chanID),
			token:  token,
			status: http.StatusOK,
		},
		"read page with negative offset": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=-1&limit=10", ts.URL, chanID),
			token:  token,
			status: http.StatusBadRequest,
		},
		"read page with negative limit": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=0&limit=-10", ts.URL, chanID),
			token:  token,
			status: http.StatusBadRequest,
		},
		"read page with zero limit": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=0&limit=0", ts.URL, chanID),
			token:  token,
			status: http.StatusBadRequest,
		},
		"read page with non-integer offset": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=abc&limit=10", ts.URL, chanID),
			token:  token,
			status: http.StatusBadRequest,
		},
		"read page with non-integer limit": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=0&limit=abc", ts.URL, chanID),
			token:  token,
			status: http.StatusBadRequest,
		},
		"read page with invalid channel id": {
			url:    fmt.Sprintf("%s/channels//messages?offset=0&limit=10", ts.URL),
			token:  token,
			status: http.StatusBadRequest,
		},
		"read page with invalid token": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=0&limit=10", ts.URL, chanID),
			token:  invalid,
			status: http.StatusForbidden,
		},
		"read page with multiple offset": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=0&offset=1&limit=10", ts.URL, chanID),
			token:  token,
			status: http.StatusBadRequest,
		},
		"read page with multiple limit": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=0&limit=20&limit=10", ts.URL, chanID),
			token:  token,
			status: http.StatusBadRequest,
		},
		"read page with empty token": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=0&limit=10", ts.URL, chanID),
			token:  "",
			status: http.StatusForbidden,
		},
		"read page with default offset": {
			url:    fmt.Sprintf("%s/channels/%s/messages?limit=10", ts.URL, chanID),
			token:  token,
			status: http.StatusOK,
		},
		"read page with default limit": {
			url:    fmt.Sprintf("%s/channels/%s/messages?offset=0", ts.URL, chanID),
			token:  token,
			status: http.StatusOK,
		},
	}

	for desc, tc := range cases {
		req := testRequest{
			client: ts.Client(),
			method: http.MethodGet,
			url:    tc.url,
			token:  tc.token,
		}
		res, err := req.make()
		assert.Nil(t, err, fmt.Sprintf("%s: unexpected error %s", desc, err))
		assert.Equal(t, tc.status, res.StatusCode, fmt.Sprintf("%s: expected %d got %d", desc, tc.status, res.StatusCode))
	}
}
