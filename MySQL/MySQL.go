package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
    go操作MySQL
	go get github.com/go-sql-driver/mysql
*/
func main() {
	// arg1：驱动名 arg2：数据源
	db, err := sql.Open("mysql", "root:SwellTestMySQL@(47.101.50.101:3306)/swell-dev")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	//insert(db)
	//update(db)
	//deleteData(db)
	//prepare(db)
	//singleQuery(db)
	multiQuery(db)
}

// 多行查询
func multiQuery(db *sql.DB) {
	var id, age int
	var name string
	rows, err := db.Query("SELECT id, `name`, age FROM test")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("id =", id, "name =", name, "age =", age)
	}
}

// 单行查询
func singleQuery(db *sql.DB) {
	var id, age int
	var name string
	row := db.QueryRow("SELECT id, `name`, age FROM test WHERE id = 1")
	err := row.Scan(&id, &name, &age)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("id =", id, "name =", name, "age =", age)
}

func insert(db *sql.DB) {
	insertSql := "INSERT INTO test(id, name, age) VALUES(2, 'test2', 25)"
	result, err := db.Exec(insertSql)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", result)
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result.LastInsertId() =", lastInsertId)

	// RowsAffected：受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result.RowsAffected() =", rowsAffected)
}

func update(db *sql.DB) {
	updateSql := "UPDATE test SET `name` = 'name1' WHERE id = 1"
	result, err := db.Exec(updateSql)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", result)
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result.LastInsertId() =", lastInsertId)

	// RowsAffected：受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result.RowsAffected() =", rowsAffected)
}

func deleteData(db *sql.DB) {
	delSql := "DELETE FROM test WHERE id = 2"
	result, err := db.Exec(delSql)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", result)
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result.LastInsertId() =", lastInsertId)

	// RowsAffected：受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result.RowsAffected() =", rowsAffected)
}

// 预处理
func prepare(db *sql.DB) {
	params := [2][3]string{{"2", "name3", "25"}, {"3", "name4", "30"}}
	stmt, err := db.Prepare("INSERT INTO test(id, `name`, age) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range params {
		stmt.Exec(p[0], p[1], p[2])
	}
}
