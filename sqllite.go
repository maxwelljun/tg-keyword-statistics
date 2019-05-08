package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

var db = &sql.DB{}
var dbnew, dbinc, dbhas *sql.Stmt

func dbNewKeyword(word string) {
	//默认添加新的关键词时,统计数量是1,所以不需要单独进行加1
	_, err := dbnew.Exec(word)
	checkErr(err)
}

func dbKeywordInc(word string) {
	_, err := dbinc.Exec(word)
	checkErr(err)
}

func dbKeyword0(unit string) {
	_, err := db.Exec("update statistics set " + unit + "count=0 where 1=1")
	checkErr(err)
}

func dbHasKeyword(word string) bool {
	rows, _ := dbhas.Query(word)
	if rows.Next() {
		rows.Close()
		return true
	} else {
		rows.Close()
		return false
	}
}

type kc struct {
	count   int
	keyword string
}

func dbFetchTopKeyword(unit string, limit int) []kc {
	rows, err := db.Query("SELECT keyword," + unit + "count FROM main.statistics WHERE " + unit + "count>0 order by " + unit + "count desc limit 0," + strconv.Itoa(limit))
	checkErr(err)
	var count int
	var keyword string
	kcs := []kc{}
	for rows.Next() {
		err = rows.Scan(&keyword, &count)
		checkErr(err)
		kcs = append(kcs, kc{count: count, keyword: keyword})
	}
	rows.Close()
	return kcs
}

func dbCountQuery(unit string) int {
	rows, err := db.Query("SELECT sum(" + unit + "count) as sumquery FROM main.statistics WHERE " + unit + "count>0")
	checkErr(err)
	var sumquery int
	if rows.Next() {
		err = rows.Scan(&sumquery)
		if err != nil {
			rows.Close()
			return 0
		}
	}
	rows.Close()
	return sumquery
}

func dbread() {
	//	rows, err := db.Query("SELECT * FROM keyword_reply")
	//	checkErr(err)
	//	var groupid int64
	//	var kvjson string
	//
	//	for rows.Next() {
	//		err = rows.Scan(&groupid, &kvjson)
	//		checkErr(err)
	//		kvs := json2kvs(kvjson)
	//		allkvs[groupid] = kvs
	//		groups = append(groups, groupid)
	//	}
	//	rows.Close()
	//
	rows, err := db.Query("SELECT value FROM settings")
	checkErr(err)
	var value string
	rows.Next()
	err = rows.Scan(&value)
	checkErr(err)
	TELEGRAM_TOKEN = value
	rows.Close()
}

func dbopen() {
	dbb, err := sql.Open("sqlite3", "./bot.db")
	checkErr(err)
	db = dbb

	dbnew, err = db.Prepare("INSERT INTO statistics(keyword) values(?)")
	checkErr(err)
	dbhas, err = db.Prepare("SELECT sumcount from statistics where keyword=?")
	checkErr(err)
	dbinc, err = db.Prepare(
		"update statistics set sumcount=sumcount+1, yearcount=yearcount+1, " +
			"monthcount=monthcount+1, weekcount=weekcount+1, daycount=daycount+1, hourcount=hourcount+1 where keyword=?")
	checkErr(err)

}

func dbclose() {
	db.Close()
}

func dbinit(token string) {
	_dbcreate()
	stmt, err := db.Prepare("INSERT INTO settings(key, value) values(?,?)")
	checkErr(err)
	_, err = stmt.Exec("token", token)
	checkErr(err)
}

func _dbcreate() {
	settings := `
    CREATE TABLE IF NOT EXISTS settings(
        key VARCHAR(64) PRIMARY KEY,
        value VARCHAR(200) NULL 
    );
    `

	statistics := `
    CREATE TABLE IF NOT EXISTS statistics(
		keyword VARCHAR(20) PRIMARY KEY,
		sumcount INTEGER default 1,
		yearcount INTEGER default 1,
		monthcount INTEGER default 1,
		weekcount INTEGER default 1,
		daycount INTEGER default 1,
		hourcount INTEGER default 1
    );
    `

	db.Exec(settings)
	db.Exec(statistics)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
