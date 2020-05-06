package cache

import (
	"database/sql"
	"log"
	"os"
	"sync"

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

var once sync.Once

// ResultDbPath 保存结果数据库
const ResultDbPath string = "D:/Project/QuickFinder/result.db"

func init() {
	sql.Register("sqlite3_conn", &sqlite.SQLiteDriver{})
}

// Context 数据库缓存
type Context struct {
	db     *sql.DB
	DbPath string
}

// ConnectDB 连接数据库
func ConnectDB(dbPath string) (*Context, error) {
	db, err := sql.Open("sqlite3_conn", dbPath)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &Context{db, dbPath}, nil
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
func CreateDB() error {
	var sqlString = "CREATE TABLE IF NOT EXISTS search_result (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"name VARCHAR(255)," +
		"result_type VARCHAR(255)," +
		"modified VARCHAR(255)," +
		"full_path TEXT" +
		")"
	if DBIsExisted(ResultDbPath) {
		return nil
	}
	c, err := ConnectDB(ResultDbPath)
	if err != nil {
		log.Fatal(err)
	}
	statement, err := c.db.Prepare(sqlString)
	_, _ = statement.Exec()
	defer c.db.Close()
	return nil
}

// InsertResult 插入一个搜索结果
func InsertResult(item *model.ResultItem) error {
	var sqlString = "INSERT INTO search_result (" +
		"name" +
		"result_type" +
		"modified" +
		"full_path" +
		")" +
		"VALUES" +
		"(?, ?, ?, ?)"
	var c *Context
	if !DBIsExisted(ResultDbPath) {
		CreateDB()
	}
	c, _ = ConnectDB(ResultDbPath)
	statement, _ := c.db.Prepare(sqlString)
	_, _ = statement.Exec(item.Name, item.ResultType, item.Modified, item.FullPath)
	defer c.db.Close()
	return nil
}

// QueryResult 搜索结果
func QueryResult(searchPattern string) error {
	if !DBIsExisted(ResultDbPath) {
		return nil
	}
	var sqlString = "SELECT id, name" +
		"FROM search_result"
	c, _ := ConnectDB(ResultDbPath)
	_, _ = c.db.Query(sqlString)
	defer c.db.Close()
	return nil
}

// DeleteDB 删除数据库
func DeleteDB() error {
	_ = os.Remove(ResultDbPath)
	return nil
}
