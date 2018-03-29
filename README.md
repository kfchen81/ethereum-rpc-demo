
ethereum-rpc-demo demonstrate how to write your own RPC server base on the rpc framework provided by go-ethereum

# Dependency
install: `go get github.com/ethereum/go-ethereum`

# Run
run: `go run main.go`

if server start successfully, you will get:
```
RPC HTTP endpoint opened url http://127.0.0.1:9010 cors  vhosts
```

our service is registered in rpc server under the namespace `kfchen`

# Test
use curl to post data to RPC server

### kfchen_echo
```
curl -H "Content-Type: application/json" -X POST --data '{"jsonrpc":"2.0","method":"kfchen_echo","params":["ping"],"id":67}' http://127.0.0.1:9010
```

Result:
```
{"jsonrpc":"2.0","id":67,"result":"ping from echo"}
```

### kfchen_getPerson
```
curl -H "Content-Type: application/json" -X POST --data '{"jsonrpc":"2.0","method":"kfchen_getPerson","params":[{"name":"robert"}],"id":67}' http://127.0.0.1:9010
```

Result:
```
{"jsonrpc":"2.0","id":67,"result":{"desc":"good man","name":"kfchen"}}
```

### rpc_modules
Namespace `rpc` is registered by ethereum rpc server itself when it started. We can use it to retrieve some meta info about rpc server itself.

```
curl -H "Content-Type: application/json" -X POST --data '{"jsonrpc":"2.0","method":"rpc_modules","params":[],"id":67}' http://127.0.0.1:9010
```

Result:
```
{"jsonrpc":"2.0","id":67,"result":{"kfchen":"1.0","rpc":"1.0"}}
```

We can see `kfchen namespace` and `rpc namespace`

# See Also
[Ethereum JSON-RPC Wiki](https://github.com/ethereum/wiki/wiki/JSON-RPC)