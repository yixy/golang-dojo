package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func queryCount() (count int, err error) {
	//根因错误，需要进行wrap返回
	return count,errors.Wrap(sql.ErrNoRows, "query count error")
}

func main() {
	count, err := queryCount()
	if err != nil {
		fmt.Printf("%+v",err)
		return
	}
	fmt.Println(count)
}
