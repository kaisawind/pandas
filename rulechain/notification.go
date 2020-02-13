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

package rulechain

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RuleChainNotification RuleChainNotification
type RuleChainNotification struct {
	UserID      string `json:"userID"`
	RuleChainID string `json:"rulechainID"`
}

// Validate validates this deployment
func (m *RuleChainNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RuleChainNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleChainNotification) UnmarshalBinary(b []byte) error {
	var res RuleChainNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}