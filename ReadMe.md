# Smart Bridge Ark

### Getting Started

Installation


Per-Install
* golang 1.9 above.
  * [dep](https://github.com/golang/dep)
  * [realize](https://github.com/tockins/realize) optional
* node v9.4
  * yarn module
  * webpack
* docker or cockroachdb node
  * if you have docker, you setup docker with `startdb.sh`
  * `docker exec -it mycockroahdb ./cockroach sql --insecure`
  * `CREATE DATABASE ark_dev;`
  * ^D
  * cd $GOPATH/src/github/masato25/smart_bridge_ark/app/setup
  * `go test` -> this will create & reset all database tables for you.

How to building

* cd $GOPATH/src/github/masato25/smart_bridge_ark
* `dep ensure`
* `go get ./...`
* `yarn`
* `npm run build`
* `go run main.go`
