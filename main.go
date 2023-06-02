package main

import (
	// "encoding/json"
	"fmt"
	"time"
	// "github.com/cloudflare/cfssl/cli/version"
	// "github.com/syndtr/goleveldb/leveldb"
)

type test struct {
	blocknum    int
	txnlist     string
	blockstatus string
	timestamp   time.Time
	prevhash    []byte
	currenthash []byte
}
type trasaction struct {
	value   string
	version string
	valid   bool
}
type blockstatus string

const (
	committed blockstatus = "committed"
	pending   blockstatus = "pending"
)

func (b blockstatus) String() string {
	return string(b)
}

type t1 interface{
	pushvalidtxn()
	updateblockstatus()

}

func (t test)pushvalidtxn(){

}
func (t test) updateblockstatus(){

}
func level (){
	// db, err := leveldb.OpenFile("path/to/db", nil)
}


func main() {


	str := committed
	fmt.Println(str)
}
