package db

import _ "github.com/go-sql-driver/mysql"
import (
	"database/sql"
	"fmt"
)

func connetdb() (b *sql.DB) {
	b, err := sql.Open("mysql", "root:root@/wechat")
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		return b
	}

}
func Exec(sqlstr string) error {
	db := connetdb()
	_, err := db.Exec(sqlstr)
	if err != nil {
		fmt.Println(err)

	}
	return err

}
