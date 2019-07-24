// Copyright (c) 2014-2017 The btcsuite developers
// Copyright (c) 2019 The Namecoin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package nmcrpcclient

import (
	"github.com/btcsuite/btcd/rpcclient"
)

// Client represents a Namecoin RPC client which allows easy access to the
// various RPC methods available on a Namecoin RPC server.  Each of the wrapper
// functions handle the details of converting the passed and return types to and
// from the underlying JSON types which are required for the JSON-RPC
// invocations
//
// The client provides each RPC in both synchronous (blocking) and asynchronous
// (non-blocking) forms.  The asynchronous forms are based on the concept of
// futures where they return an instance of a type that promises to deliver the
// result of the invocation at some future time.  Invoking the Receive method on
// the returned future will block until the result is available if it's not
// already.
type Client struct {
	*rpcclient.Client
}

// New creates a new RPC client based on the provided connection configuration
// details.  The notification handlers parameter may be nil if you are not
// interested in receiving notifications and will be ignored if the
// configuration is set to run in HTTP POST mode.
func New(config *rpcclient.ConnConfig, ntfnHandlers *rpcclient.NotificationHandlers) (*Client, error) {
	btcClient, err := rpcclient.New(config, ntfnHandlers)
	if err != nil {
		return nil, err
	}

	return &Client{btcClient}, nil
}
