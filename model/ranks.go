package model

import (
	"time"

	"github.com/rikimaru-tokyo/gin_sqlx_lesson02/myDatabase"
)

//データの受け手となる構造体を宣言
//`db`タグは必須。
//タグの値はDBテーブルのカラム名に合わせること。
type RankResult struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

// InsertRank() ranksテーブルにバルクインサートを実行する。
func InsertRank() (bool, error) {
	//データベース接続
	db := myDatabase.DbInit()
	now := time.Now()

	sql := `INSERT INTO ranks 
				(id, name, created_at, updated_at)
			VALUES
				(:id, :name, :cr, :up);`

	//SQLで指定されているパラメータと、連想配列の要素をもれなく統一させること。
	inserts := []map[string]interface{}{
		{"id": 4, "name": "Platinum", "cr": now, "up": now},
		{"id": 5, "name": "Diamond", "cr": now, "up": now},
	}

	// transactionオブジェクトを元にSQLを実行する。
	rows, err := db.NamedExec(sql, inserts)
	if err != nil {
		return false, err
	}

	// INSERT件数カウント
	r, err := rows.RowsAffected()
	if err != nil || r < 1 {
		return false, err
	}

	return true, nil

}

// DeleteRank() ranksテーブルからデータを2件削除する。トランザクションを用いる。
func DeleteRank() (bool, error) {
	//データベース接続
	db := myDatabase.DbInit()

	sql := `DELETE FROM ranks WHERE name = :del1 OR name = :del2;`

	deletes := map[string]interface{}{
		"del1": "Platinum",
		"del2": "Diamond",
	}

	// トランザクション開始
	transaction, err := db.Beginx()
	if err != nil {
		return false, err
	}

	// transactionオブジェクトを元にSQLを実行する。
	rows, err := transaction.NamedExec(sql, deletes)
	if err != nil {
		//実行に失敗した場合はロールバックさせ、トランザクションを終了させる。
		_ = transaction.Rollback()
		return false, err
	}

	// 成功した場合はコミットさせ、トランザクションを終了させる。
	if err := transaction.Commit(); err != nil {
		return false, err
	}

	// DELETE件数カウント
	r, err := rows.RowsAffected()
	if err != nil || r < 1 {
		return false, err
	}

	return true, nil

}
