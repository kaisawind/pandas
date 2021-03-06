// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"context"

	"github.com/cloustone/pandas/mainflux"
	"github.com/cloustone/pandas/users"
	"google.golang.org/grpc"
)

var _ mainflux.AuthNServiceClient = (*serviceMock)(nil)

type serviceMock struct {
	users map[string]string
}

// NewUsersService creates mock of users service.
func NewUsersService(users map[string]string) mainflux.AuthNServiceClient {
	return &serviceMock{users}
}

func (svc serviceMock) Identify(ctx context.Context, in *mainflux.Token, opts ...grpc.CallOption) (*mainflux.UserID, error) {
	if id, ok := svc.users[in.Value]; ok {
		return &mainflux.UserID{Value: id}, nil
	}
	return nil, users.ErrUnauthorizedAccess
}

func (svc serviceMock) Issue(ctx context.Context, in *mainflux.IssueReq, opts ...grpc.CallOption) (*mainflux.Token, error) {
	if id, ok := svc.users[in.GetIssuer()]; ok {
		switch in.Type {
		default:
			return &mainflux.Token{Value: id}, nil
		}
	}
	return nil, users.ErrUnauthorizedAccess
}
