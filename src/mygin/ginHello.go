package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type ActivityPv struct {
	Id           int    `db:"id"`
	ActivityId   string `db:"activity_id"`
	PositionType int    `db:"position_type"`
	DeviceType   int    `db:"device_type"`
	Ip           string `db:"ip"`
	DeviceNo     string `db:"device_no"`
}
func init() {
	database, err := sqlx.Open("mysql", "root:Nivelle123@tcp(localhost:3306)/javaguides")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		defer Db.Close()
		return
	}
	Db = database
}
var Db *sqlx.DB
func SelectActivity() []ActivityPv {
	var activities []ActivityPv

	err := Db.Select(&activities, "select id,activity_id,position_type,device_type,ip,device_no from activity_pv ")
	if err != nil {
		fmt.Println("exec failed", err)
	}
	fmt.Println("select ok:", activities)
	return activities
}


func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/myGin", func(c *gin.Context) {
		c.String(http.StatusOK, "hello myGin!,data=%s", SelectActivity())
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
