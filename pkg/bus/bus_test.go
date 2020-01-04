//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package bus

import (
	"errors"
	"fmt"
	"testing"
)

type TestQuery struct {
	Id   int64
	Resp string
}

func TestQueryHandlerReturnsError(t *testing.T) {
	bus := New()

	bus.AddHandler(func(query *TestQuery) error {
		return errors.New("handler error")
	})

	err := bus.Dispatch(&TestQuery{})

	if err == nil {
		t.Fatal("Send query failed " + err.Error())
	} else {
		t.Log("Handler error received ok")
	}
}

func TestQueryHandlerReturn(t *testing.T) {
	bus := New()

	bus.AddHandler(func(q *TestQuery) error {
		q.Resp = "hello from handler"
		return nil
	})

	query := &TestQuery{}
	err := bus.Dispatch(query)

	if err != nil {
		t.Fatal("Send query failed " + err.Error())
	} else if query.Resp != "hello from handler" {
		t.Fatal("Failed to get response from handler")
	}
}

func TestEventListeners(t *testing.T) {
	bus := New()
	count := 0

	bus.AddEventListener(func(query *TestQuery) error {
		count += 1
		return nil
	})

	bus.AddEventListener(func(query *TestQuery) error {
		count += 10
		return nil
	})

	err := bus.Publish(&TestQuery{})

	if err != nil {
		t.Fatal("Publish event failed " + err.Error())
	} else if count != 11 {
		t.Fatal(fmt.Sprintf("Publish event failed, listeners called: %v, expected: %v", count, 11))
	}
}