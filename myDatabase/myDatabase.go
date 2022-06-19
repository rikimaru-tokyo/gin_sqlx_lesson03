package myDatabase

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DbInit() データベースに接続する。
func DbInit() *sqlx.DB {

	//dsn形式：username:password@protocol(address)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&loc=Asia%%2fTokyo",
		os.Getenv("MYSQL_USER"), 
        os.Getenv("MYSQL_PASSWORD"), 
        os.Getenv("MYSQL_HOST"), 
        os.Getenv("MYSQL_DATABASE"),
	)

	//MySQLに接続
    //sqlx.Openの第1引数は固定値でDBの種類を入力(今回は「mysql」)、第2引数は固定値でDSNを入れる。
	conn, err := sqlx.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

    //接続オブジェクトのポインタを返します。
	return conn
}

