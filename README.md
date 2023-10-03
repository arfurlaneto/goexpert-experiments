# golang-grcp-example

## Compiling Protobuf

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go-grpc_out=. proto/entity.proto
```

## Connect with Evans

```bash
evans -r repl
```
## Evans commands


```bash
show package
package <package_name>
show message
desc <message_name>
show service
service <service_name>
call <service_rpc_name>
ctrl + d
```
