package cache

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	sqlite "github.com/mattn/go-sqlite3"

	"github.com/nashera/QuickFinder/model"
)

// func build() {
// 	database, _ := sql.Open("sqlite3", "./nraboy.db")
// 	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS search_result (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
// 	_, _ = statement.Exec()
// 	statement, _ = database.Prepare("INSERT INTO search_result (firstname, lastname) VALUES (?, ?)")
// 	_, _ = statement.Exec("Nic", "Raboy")
// 	rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
// 	var id int
// 	var firstname string
// 	var lastname string
// 	for rows.Next() {
// 		_ = rows.Scan(&id, &firstname, &lastname)
// 		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
// 	}
// }

// Context 数据库缓存
type Context struct {
	db     *sql.DB
	dbPath string
}

// ConnectDB 连接数据库
func ConnectDB(dbPath string) (*Context, string) {
	sql.Register("sqlite3_conn", &sqlite.SQLiteDriver{})
	db, err := sql.Open("sqlite3_conn", dbPath)
	if err != nil {
		return nil, err.Error()
	}
	if err = db.Ping(); err != nil {
		return nil, err.Error()
	}
	return &Context{db, dbPath}, ""
}

// DBIsExisted 判断数据库是否存在
func DBIsExisted(dbPath string) bool {
	existed := true
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

// CreateDB 创建sqlite3数据库， 用于缓存
func (c *Context) CreateDB() error {
	var sqlString = "CREATE TABLE IF NOT EXISTS search_result (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"name VARCHAR(255)," +
		"result_type VARCHAR(255)," +
		"modified VARCHAR(255)," +
		"full_path TEXT" +
		")"
	fmt.Println(sqlString)
	statement, err := c.db.Prepare(sqlString)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = statement.Exec()
	return nil
}

// InsertResult 插入一个搜索结果
func (c *Context) InsertResult(item *model.ResultItem) error {
	var sqlString = "INSERT INTO search_result (" +
		"name" +
		"result_type" +
		"modified" +
		"full_path" +
		")" +
		"VALUES" +
		"(?, ?, ?, ?)"
	statement, _ := c.db.Prepare(sqlString)
	_, _ = statement.Exec(item.Name, item.ResultType, item.Modified, item.FullPath)

	return nil
}

// QueryResult 搜索结果
func (c *Context) QueryResult(searchPattern string) error {
	var sqlString = "SELECT id, name" +
		"FROM search_result"
	_, _ = c.db.Query(sqlString)
	return nil
}

// DeleteDB 删除数据库
func (c *Context) DeleteDB() error {
	_ = os.Remove(c.dbPath)
	return nil
}
