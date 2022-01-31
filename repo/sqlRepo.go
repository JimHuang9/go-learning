package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //為什麼要用底線_來占著位子？
)

type Teacher struct {
	Id   int
	Name string
	Age  int
}

func GetTeachers() []Teacher {
	techers := []Teacher{}
	rows, err := db.Query("SELECT `Id`, `Name`,`Age` FROM  `School`.`Teacher`")
	for rows.Next() {
		var tId int
		var tName string
		var tAge int
		err = rows.Scan(&tId, &tName, &tAge)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%q %d\n", tName, tAge)
		var techer = Teacher{
			tId,
			tName,
			tAge,
		}
		techers = append(techers, techer)
	}
	defer rows.Close()
	return techers
}

func InserTeacher(name string, age int) int64 {
	rs, err := db.Exec("INSERT INTO `School`.`Teacher` (`Name`,`Age`) VALUES (?,?)", name, age)
	if err != nil {
		log.Println(err)
	}
	rowCount, err := rs.RowsAffected()
	rowId, err := rs.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("新增 %d 筆資料，id = %d \n", rowCount, rowId)
	return rowId
}

func UpdateTeacher(name string, age int, id int64) {
	rs, err := db.Exec("UPDATE  `School`.`Teacher` SET `NAME`=? ,`Age`=? WHERE `Id`=? ", name, age, id)
	if err != nil {
		log.Println(err)
	}
	rowCount, err := rs.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("更新 %d 筆資料", rowCount)
}

func DeleteTeacher(id int64) {
	rs, err := db.Exec("Delete FROM  `School`.`Teacher` WHERE `Id`=? ", id)
	if err != nil {
		log.Println(err)
	}
	rowCount, err := rs.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("刪除 %d 筆資料", rowCount)
}

var db *sql.DB

// 與DB連線。 init() 初始化，時間點比 main() 更早。
func init() {
	dbConnect, err := sql.Open(
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/",
	)

	if err != nil {
		log.Fatalln(err)
	}

	err = dbConnect.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	db = dbConnect // 用全域變數接

	db.SetMaxOpenConns(10) // 可設置最大DB連線數，設<=0則無上限（連線分成 in-Use正在執行任務 及 idle執行完成後的閒置 兩種）
	db.SetMaxIdleConns(10) // 設置最大idle閒置連線數。
	// 更多用法可以 進 sql.DBStats{}、sql.DB{} 裡面看
}
