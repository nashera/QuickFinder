package cache

import (
	"database/sql"
	"fmt"
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

// DBContext 数据库缓存
type DBContext struct {
	db     *sql.DB
	DbPath string
}

// ConnectDB 连接数据库
func ConnectDB(dbPath string) (*DBContext, error) {
	db, err := sql.Open("sqlite3_conn", dbPath)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DBContext{db, dbPath}, nil
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
func (dc *DBContext) CreateDB() error {
	var sqlString = `
	CREATE TABLE IF NOT EXISTS search_result
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		result_name VARCHAR(255),
		result_type VARCHAR(255),
		modified DATETIME,
		full_path TEXT
	);
	CREATE TABLE IF NOT EXISTS sample
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sample_name VARCHAR(255),
		lib_name VARCHAR(255),
		doctor VARCHAR(255),
		hospital VARCHAR(255),
		result_id INTEGER,
		foreign key(result_id) references search_result(id)
	);
	`
	_, err := dc.db.Exec(sqlString)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// InsertResult 插入一个搜索结果
func (dc *DBContext) InsertResult(item *model.ResultItem) error {
	var insertResultSQLString = `
	INSERT INTO search_result (
		result_name,
		result_type,
		modified,
		full_path
	)
	VALUES
	(?, ?, ?, ?);
	`
	// fmt.Println(sqlString)
	// fmt.Println(item.Modified)
	if !DBIsExisted(dc.DbPath) {
		dc.CreateDB()
	}
	statement, _ := dc.db.Prepare(insertResultSQLString)
	insertResult, err := statement.Exec(item.Name, item.ResultType, item.Modified.Unix(), item.FullPath)
	if err != nil {
		fmt.Printf("insert error: %v", err)

	}
	lastID, err := insertResult.LastInsertId()
	// fmt.Println(lastID)
	// _, err = insertResult.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	// _, _ = c.db.Exec(sqlString, item.Name)
	var insertSampleSQLString = `
	INSERT INTO sample (
		sample_name,
		lib_name,
		doctor,
		hospital,
		result_id
	)
	VALUES
	(?, ?, ?, ?, ?);
	`
	statement, _ = dc.db.Prepare(insertSampleSQLString)
	for _, sample := range item.Samples {
		_, err := statement.Exec(sample.SampleName, sample.LibName, sample.Doctor, sample.Hospital, lastID)
		if err != nil {
			log.Fatal(err)
		}

	}
	return nil
}

// QueryResult 搜索结果
func (dc *DBContext) QueryResult(query string) []*model.ResultItem {
	if !DBIsExisted(dc.DbPath) {
		return nil
	}
	var sqlString string
	sqlString = `
	SELECT 
	search_result.id, 
	search_result.result_name, 
	search_result.result_type,
	search_result.modified,
	search_result.full_path
	FROM 
	search_result INNER JOIN sample ON search_result.id = sample.result_id 
	WHERE 
	search_result.result_name LIKE ?
	OR search_result.full_path LIKE ?
	OR sample.doctor LIKE ?
	OR sample.hospital LIKE ?
	`
	rows, err := dc.db.Query(sqlString, "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var results []*model.ResultItem
	for rows.Next() {
		resultitem := model.ResultItem{}
		var resultID int32
		err = rows.Scan(&resultID, &resultitem.Name, &resultitem.ResultType, &resultitem.Modified, &resultitem.FullPath)
		if err != nil {
			panic(err)
		}
		resultitem.Samples = dc.QuerySampleByResultID(resultID)
		results = append(results, &resultitem)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return results
}

// QuerySample 搜索结果
func (dc *DBContext) QuerySample(query string) []*model.Sample {
	if !DBIsExisted(dc.DbPath) {
		return nil
	}
	var sqlString string
	sqlString = `
	SELECT 
	sample_name,
	lib_name,
	doctor,
	hospital
	FROM 
	sample WHERE 
	doctor LIKE ?
	OR hospital LIKE ?
	`
	rows, err := dc.db.Query(sqlString, "%"+query+"%", "%"+query+"%")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var samples []*model.Sample
	for rows.Next() {
		sample := model.Sample{}
		err = rows.Scan(&sample.SampleName, &sample.LibName, &sample.Doctor, &sample.Hospital)
		if err != nil {
			panic(err)
		}
		samples = append(samples, &sample)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return samples
}

// DeleteDB 删除数据库
func DeleteDB(dbPath string) error {
	_ = os.Remove(dbPath)
	return nil
}

// QuerySampleByResultID 根据result_id搜索Sample
func (dc *DBContext) QuerySampleByResultID(resultID int32) []*model.Sample {
	if !DBIsExisted(dc.DbPath) {
		return nil
	}
	var sqlString string
	sqlString = `
	SELECT 
	sample_name,
	lib_name,
	doctor,
	hospital
	FROM 
	sample WHERE 
	result_id = ?
	`
	rows, err := dc.db.Query(sqlString, resultID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var samples []*model.Sample
	for rows.Next() {
		sample := model.Sample{}
		err = rows.Scan(&sample.SampleName, &sample.LibName, &sample.Doctor, &sample.Hospital)
		if err != nil {
			panic(err)
		}
		samples = append(samples, &sample)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return samples

}

//CloseDB close DB connect
func (dc *DBContext) CloseDB() {
	dc.db.Close()
}
