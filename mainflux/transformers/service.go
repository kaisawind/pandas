// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package transformers

import "github.com/cloustone/pandas/mainflux/broker"

// Transformer specifies API form Message transformer.
type Transformer interface {
	// Transform Mainflux message to any other format.
	Transform(broker.Message) (interface{}, error)
}
