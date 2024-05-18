package main

import (
	"context"
	"fmt"

	"github.com/jasonz3g/gen/examples/biz"
	"github.com/jasonz3g/gen/examples/conf"
	"github.com/jasonz3g/gen/examples/dal"
	"github.com/jasonz3g/gen/examples/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Query(context.Background())
}
