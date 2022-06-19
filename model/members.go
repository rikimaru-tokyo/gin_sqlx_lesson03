package model

import (
	"github.com/rikimaru-tokyo/gin_sqlx_lesson02/myDatabase"
)

//データの受け手となる構造体を宣言
type MembersResult struct {
	MemberID   int    `db:"id"`
	MemberName string `db:"name"`
}

// UseSelect Selectメソッドを使ったデータ取得
func UseSelect(memberId int) ([]MembersResult, error) {
	//データベース接続
	db := myDatabase.DbInit()

	sql := `SELECT id, name FROM members WHERE id >= ?;`

	var member []MembersResult

	// Select()は複数行のデータ取得時に使用する。
	err := db.Select(&member, sql, memberId)
	if err != nil {
		return nil, err
	}

	return member, nil

}

// UseGet Getメソッドを使ったデータ取得
func UseGet(memberId int) (bool, error) {
	//データベース接続
	db := myDatabase.DbInit()

	sql := `SELECT EXISTS (SELECT id FROM members WHERE id = ?);`

	var exists bool

	// Get()は単数データ取得時に使用する。
	// データが見つからない場合はエラーを出すので注意。
	// 「SELECT COUNT」や「SELECT EXISTS」のケースに使いやすい。
	err := db.Get(&exists, sql, memberId)
	if err != nil {
		return false, err
	}

	return exists, nil

}
