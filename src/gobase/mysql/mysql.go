package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type activityPv struct {
	id         int    `db:id`
	activityId string `db:activity_id`

	positionType int    `db:position_type`
	deviceType   int    `db:device_type`
	ip           string `db:ip`
	deviceNo     string `db:device_no`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:Nivelle123@tcp(localhost:3306)/javaguides")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		defer Db.Close()
		return
	}
	Db = database
}

func main() {
	r, err := Db.Exec("insert into activity_pv(activity_id, position_type, device_type,ip,device_no)values(?, ?, ?,?,?)", "go1", "1", "2","192.168.7.89","go2")
	if err != nil {
		fmt.Println("exec failed1, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed2, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}
