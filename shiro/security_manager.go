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
package shiro

import (
	"errors"

	. "github.com/cloustone/padnas/shiro/realms"
	"github.com/cloustone/pandas/shiro/adaptors"
	"github.com/cloustone/pandas/shiro/options"
	"github.com/cloustone/pandas/shiro/realms"
)

// SecurityManager is responsible for authenticate and simple authorization
type SecurityManager interface {
	UseAdaptor(adaptors.Adaptor)
	Authenticate(principal *Principal) error
	Authorize(principal Principal, subject *Subject, action string) error
	GetAuthzDefinitions(principal principal) ([]*AuthzDefinition, error)
	GetPrincipalDefinition(principal Principal) (*PrincipalDefinition, error)
	GetPrincipalAllowableSubjects(principal Principal) ([]*Subject, error)
}

// NewSecurityManager create security manager to hold all realms for
// authenticate
func NewSecurityManager(servingOptions *options.ServingOptions) SecurityManager {
	return nil
}

// defaultSecuriityManager
type defaultSecurityManager struct {
	servingOptions *options.ServingOptions
	realms         []realms.Realm
}

// newDefaultSecurityManager return security manager instance
// All realms are created here, if failed, shiro must be restarted
func newDefaultSecurityManager(servingOptions *options.ServingOptions) *defaultSecurityManager {
	realmOptions := NewRealmOptions(servingOptions.RealmConfigFile)
	realms := []Ream{}

	for _, options := range realmOptions {
		realms = append(realms, NewRealm(option))
	}
	return &defaultSecurityManager{
		servingOptions: servingOptions,
		realms:         realms,
	}
}

// Authenticate iterate all realm to authenticate the principal
func (s *defaultSecuirytManager) Authenticate(principal *Principal) error {
	for _, realm := range s.realms {
		if err := realm.Authenticate(principal); err == nil {
			return nil
		}
	}
	return errors.New("no valid realms")
}
