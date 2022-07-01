package model

import (
	"github.com/rikimaru-tokyo/gin_sqlx_lesson02/myDatabase"
)

// INSERT、UPDATEに値を入れるための構造体
type HogeTableIn struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Birthday string `db:"birthday"`
}

func Insert() error {

	// 事前に構造体に値を詰めておく。
	in := HogeTableIn{
		ID:       1,
		Name:     "alpha",
		Birthday: "2022-01-23",
	}

	// INSERT SQLを定義する。プレースフォルダの変数名は構造体のタグ`db:"***"`で定義した名称と同じにする。
	sql := `INSERT INTO hoge_table (id, name, birthday) VALUES (:id, :name, :birthday);`
	db := myDatabase.DbInit()
	_, err := db.NamedExec(sql, in)

	if err != nil {
		return err
	}

	return nil
}

func Update() error {

	// 事前に構造体に値を詰めておく。
	in := HogeTableIn{
		ID:   1,
		Name: "bravo",
	}

	// UPDATE SQLを定義する。プレースフォルダの変数名は構造体のタグ`db:"***"`で定義した名称と同じにする。
	sql := `UPDATE hoge_table SET name = :name WHERE id = :id;`

	db := myDatabase.DbInit()
	_, err := db.NamedExec(sql, in)

	if err != nil {
		return err
	}

	return nil
}

func BulkInsert() error {

	//バルクインサートを試す。
	in := []HogeTableIn{
		{
			ID:       2,
			Name:     "charlie",
			Birthday: "2022-04-05",
		},
		{
			ID:       3,
			Name:     "delta",
			Birthday: "2022-06-07",
		},
	}

	sql := `INSERT INTO hoge_table (id, name, birthday) VALUES (:id, :name, :birthday);`

	db := myDatabase.DbInit()
	_, err := db.NamedExec(sql, in)

	if err != nil {
		return err
	}

	return nil
}

// 反映確認用
func GetAll() ([]HogeTableIn, error) {
	var data []HogeTableIn
	db := myDatabase.DbInit()
	err := db.Select(&data, "SELECT * FROM hoge_table")

	if err != nil {
		return nil, err
	}

	return data, nil

}
