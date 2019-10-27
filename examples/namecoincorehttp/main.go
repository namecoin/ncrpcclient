// Copyright (c) 2014-2017 The btcsuite developers
// Copyright (c) 2019 The Namecoin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/namecoin/btcd/rpcclient"
	"github.com/namecoin/ncrpcclient"
)

func main() {
	// Connect to local namecoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:8336",
		User:         "yourrpcuser",
		Pass:         "yourrpcpass",
		HTTPPostMode: true, // Namecoin core only supports HTTP POST mode
		DisableTLS:   true, // Namecoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := ncrpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current data for the name.
	nameData, err := client.NameShow("d/domob")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Name: %s", nameData.Name)
	log.Printf("Value: %s", nameData.Value)
}
