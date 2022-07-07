package model

import (
	"time"

	"github.com/rikimaru-tokyo/gin_sqlx_lesson03/myDatabase"
)

type TargetsTable struct {
	ID       int       `db:"id"`
	Name     string    `db:"name"`
	Birthday time.Time `db:"birthday"`
}

type TargetID struct {
	ID int `json:"id" binding:"required"`
}

type TargetUpdate struct {
	Name     string `json:"name" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
}

func FindTargetAll() ([]TargetsTable, error) {
	sql := `SELECT id, name, birthday FROM targets;`

	var result []TargetsTable
	if err := myDatabase.DbInit().Select(&result, sql); err != nil {
		return []TargetsTable{}, err
	}

	return result, nil
}

func FindTargetOne(id int) ([]TargetsTable, error) {
	sql := `SELECT id, name, birthday FROM targets WHERE id = ?;`

	var result []TargetsTable
	if err := myDatabase.DbInit().Select(&result, sql, id); err != nil {
		return []TargetsTable{}, err
	}

	return result, nil
}

func InsertTargetOne(param TargetsTable) (int, error) {
	sql := `INSERT INTO targets VALUES (:id, :name, :birthday);`

	transaction, err := myDatabase.DbInit().Beginx()
	if err != nil {
		return 0, err
	}

	result, err := transaction.NamedExec(sql, param)
	if err != nil {
		_ = transaction.Rollback()
		return 0, err
	}

	n, err := result.LastInsertId()
	if err != nil {
		_ = transaction.Rollback()
		return 0, err
	}

	if err := transaction.Commit(); err != nil {
		_ = transaction.Rollback()
		return 0, err
	}

	return int(n), nil
}

func InsertTargetBulk(param []TargetsTable) error {
	sql := `INSERT INTO targets VALUES (:id, :name, :birthday);`

	transaction, err := myDatabase.DbInit().Beginx()
	if err != nil {
		return err
	}

	_, err = transaction.NamedExec(sql, param)
	if err != nil {
		_ = transaction.Rollback()
		return err
	}

	if err = transaction.Commit(); err != nil {
		_ = transaction.Rollback()
		return err
	}

	return nil
}

func UpdateTarget(id int, param TargetUpdate) error {
	sql := `UPDATE targets SET name=?, birthday=? WHERE id = ?;`

	transaction, err := myDatabase.DbInit().Beginx()
	if err != nil {
		return err
	}

	_, err = transaction.Exec(sql, param.Name, param.Birthday, id)
	if err != nil {
		_ = transaction.Rollback()
		return err
	}

	if err = transaction.Commit(); err != nil {
		_ = transaction.Rollback()
		return err
	}

	return nil
}

func DeleteTarget(id int) error {
	sql := `DELETE FROM targets WHERE id = ?;`

	transaction, err := myDatabase.DbInit().Beginx()
	if err != nil {
		return err
	}

	_, err = transaction.Exec(sql, id)
	if err != nil {
		_ = transaction.Rollback()
		return err
	}

	if err = transaction.Commit(); err != nil {
		_ = transaction.Rollback()
		return err
	}

	return nil
}
