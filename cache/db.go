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
		"modified DATETIME," +
		"full_path TEXT" +
		");"
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
		"name, " +
		"result_type, " +
		"modified, " +
		"full_path" +
		") " +
		"VALUES " +
		"(?, ?, ?, ?);"
	// fmt.Println(sqlString)
	// fmt.Println(item.Modified)
	if !DBIsExisted(ResultDbPath) {
		CreateDB()
	}
	var c *Context
	c, _ = ConnectDB(ResultDbPath)
	statement, _ := c.db.Prepare(sqlString)
	_, _ = statement.Exec(item.Name, item.ResultType, item.Modified.Unix(), item.FullPath)
	// _, _ = c.db.Exec(sqlString, item.Name)
	defer c.db.Close()
	return nil
}

// QueryResult 搜索结果
func QueryResult() []model.ResultItem {
	if !DBIsExisted(ResultDbPath) {
		return nil
	}
	var sqlString = "SELECT name, result_type, modified, full_path " +
		"FROM search_result"
	c, _ := ConnectDB(ResultDbPath)
	rows, err := c.db.Query(sqlString)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var results []model.ResultItem
	for rows.Next() {
		resultitem := model.ResultItem{}
		err = rows.Scan(&resultitem.Name, &resultitem.ResultType, &resultitem.Modified, &resultitem.FullPath)
		if err != nil {
			panic(err)
		}
		results = append(results, resultitem)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	defer c.db.Close()
	return results
}

// DeleteDB 删除数据库
func DeleteDB() error {
	_ = os.Remove(ResultDbPath)
	return nil
}
