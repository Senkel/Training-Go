package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//инициализация базы
	db, _ := sql.Open("mysql", "root:@/go-18")
	defer db.Close()
	//выборка из таблицы
	row, err := db.Query("SELECT fio FROM students WHERE id=1")
	for row.Next() {
		var fio string
		err = row.Scan(&fio)
		PanicOnErr(err)
		fmt.Println(fio)
	}
	//вставка в таблицу
	result, err := db.Exec(
		"INSERT INTO students (`fio`, `score`) VALUES (?, 0)",
		"Ivan Ivanov",
	)
	PanicOnErr(err)

	affected, err := result.RowsAffected()
	PanicOnErr(err)
	lastID, err := result.LastInsertId()
	PanicOnErr(err)

	fmt.Println("Insert - RowsAffected", affected, "LastInsertId: ", lastID)
	//изменения записей в таблице
	res, err := db.Exec("UPDATE students set id=? WHERE fio = ?", 0, "robert senkel")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())
	//удаление записей в таблице
	r, err := db.Exec("DELETE FROM students WHERE id=?", 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.RowsAffected())

}
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
