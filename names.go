// Copyright (c) 2014-2017 The btcsuite developers
// Copyright (c) 2019 The Namecoin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ncrpcclient

import (
	"encoding/hex"
	"encoding/json"

	"github.com/btcsuite/btcd/rpcclient"

	"github.com/namecoin/ncbtcjson"
)

// *********************
// Name Lookup Functions
// *********************

// FutureNameShowResult is a future promise to deliver the result
// of a NameShowAsync RPC invocation (or an applicable error).
type FutureNameShowResult chan *rpcclient.Response

// Receive waits for the Response promised by the future and returns detailed
// information about a name.
func (r FutureNameShowResult) Receive() (*ncbtcjson.NameShowResult, error) {
	res, err := rpcclient.ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a name_show result object
	var nameShow ncbtcjson.NameShowResult

	err = json.Unmarshal(res, &nameShow)
	if err != nil {
		return nil, err
	}

	if nameShow.NameEncoding == ncbtcjson.Hex {
		var nameBytes []byte

		nameBytes, err = hex.DecodeString(nameShow.Name)
		if err != nil {
			return nil, err
		}

		nameShow.Name = string(nameBytes)
	}

	if nameShow.ValueEncoding == ncbtcjson.Hex {
		var valueBytes []byte

		valueBytes, err = hex.DecodeString(nameShow.Value)
		if err != nil {
			return nil, err
		}

		nameShow.Value = string(valueBytes)
	}

	return &nameShow, nil
}

// NameShowAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See NameShow for the blocking version and more details.
func (c *Client) NameShowAsync(name string, options *ncbtcjson.NameShowOptions) FutureNameShowResult {
	if options != nil && options.NameEncoding == ncbtcjson.Hex {
		name = hex.EncodeToString([]byte(name))
	}

	cmd := ncbtcjson.NewNameShowCmd(name, options)

	return c.SendCmd(cmd)
}

// NameShow returns detailed information about a name.
func (c *Client) NameShow(name string, options *ncbtcjson.NameShowOptions) (*ncbtcjson.NameShowResult, error) {
	return c.NameShowAsync(name, options).Receive()
}

// FutureNameScanResult is a future promise to deliver the result
// of a NameScanAsync RPC invocation (or an applicable error).
type FutureNameScanResult chan *rpcclient.Response

// Receive waits for the Response promised by the future and returns detailed
// information about a list of names.
func (r FutureNameScanResult) Receive() (ncbtcjson.NameScanResult, error) {
	res, err := rpcclient.ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a name_scan result object
	var nameScan ncbtcjson.NameScanResult
	err = json.Unmarshal(res, &nameScan)
	if err != nil {
		return nil, err
	}

	return nameScan, nil
}

// NameScanAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See NameScan for the blocking version and more details.
func (c *Client) NameScanAsync(start string, count uint32) FutureNameScanResult {
	cmd := ncbtcjson.NewNameScanCmd(start, &count, nil)
	return c.SendCmd(cmd)
}

// NameScan returns detailed information about a list of names.
func (c *Client) NameScan(start string, count uint32) (ncbtcjson.NameScanResult, error) {
	return c.NameScanAsync(start, count).Receive()
}
