package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *sql.DB
}

// 定义user结构体
type user struct {
	Id                  int16
	RESOURCE_GROUP_NAME string
	START_TIME          string
	END_TIME            string
	WATCH               string
	WATCH_TEXT          string
	ACTION              string
}

func (d *Database) Init_db(host, port string) error {
	// DSN: Data Source Name
	dsn := fmt.Sprintf("root:pass@tcp(%s:%s)/INFORMATION_SCHEMA?charset=utf8mb4&parseTime=True", host, port)

	// 打开数据库连接
	var err error
	d.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	// 检查数据库连接是否可用
	err = d.DB.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return err
	}

	return nil
}

func (d *Database) QueryRowDemo() error {
	sqlStr := "SELECT id,RESOURCE_GROUP_NAME,START_TIME,END_TIME,WATCH,WATCH_TEXT,ACTION FROM INFORMATION_SCHEMA.RUNAWAY_WATCHES "
	var u user
	// 非常重要：确保 QueryRow 之后调用 Scan 方法，否则持有的数据库链接不会被释放
	rows, err := d.DB.Query(sqlStr)
	if err != nil {
		return fmt.Errorf("query failed, err:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&u.Id, &u.RESOURCE_GROUP_NAME, &u.START_TIME, &u.END_TIME, &u.WATCH, &u.WATCH_TEXT, &u.ACTION)
		if err != nil {
			return fmt.Errorf("scan failed, err:%v", err)
		}
		//fmt.Printf("u:%#v\n", u)
		fmt.Println("当前集群 拉黑SQL信息")
		fmt.Println("RESOURCE_GROUP_NAME:", u.RESOURCE_GROUP_NAME, "START_TIME:", u.START_TIME, "END_TIME:", u.END_TIME, "WATCH:", u.WATCH, "WATCH_TEXT:", u.WATCH_TEXT, "ACTION:", u.ACTION)
		fmt.Println("id:", u.Id, "SQL:", u.WATCH_TEXT)
	}
	defer d.DB.Close()
	return nil

}

func (d *Database) Bindsql(SQL string) error {
	sqlStr := fmt.Sprintf("QUERY WATCH ADD ACTION KILL SQL TEXT EXACT TO '%s';", SQL)
	// 非常重要：确保 QueryRow 之后调用 Scan 方法，否则持有的数据库链接不会被释放
	_, err := d.DB.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("query failed, err:%v", err)
	}
	//fmt.Println("result:", result)
	//fmt.Println("err:", err)
	return nil
}

func (d *Database) BindsqlUnbind(pID int) error {
	sqlStr := fmt.Sprintf("QUERY WATCH REMOVE %d;", pID)
	fmt.Println(pID)
	fmt.Println(sqlStr)
	// 非常重要：确保 QueryRow 之后调用 Scan 方法，否则持有的数据库链接不会被释放
	_, err := d.DB.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("query failed, err:%v", err)
	}
	//fmt.Println("result:", result)
	//fmt.Println("err:", err)
	return nil
}
