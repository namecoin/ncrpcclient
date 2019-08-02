Namecoin Core HTTP POST Example
===============================

This example shows how to use the ncrpcclient package to connect to a Namecoin
Core RPC server using HTTP POST mode with TLS disabled and looks up a name.

## Running the Example

The first step is to use `go get` to download and install the ncrpcclient package:

```bash
$ go get github.com/namecoin/ncrpcclient
```

Next, modify the `main.go` source to specify the correct RPC username and
password for the RPC server:

```Go
	User: "yourrpcuser",
	Pass: "yourrpcpass",
```

Finally, navigate to the example's directory and run it with:

```bash
$ cd $GOPATH/src/github.com/namecoin/ncrpcclient/examples/namecoincorehttp
$ go run *.go
```

## License

This example is licensed under the [copyfree](http://copyfree.org) ISC License.
