package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jedib0t/go-pretty/table"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const (
	mysqlPort = 3306
	mysqlDB   = "information_schema"
)

var (
	mysqlUser     string
	mysqlPassword string
	mysqlHost     string
)

func init() {
	flag.StringVar(&mysqlUser, "u", "", "-u mysql user")
	flag.StringVar(&mysqlPassword, "p", "", "-p mysql password")
	flag.StringVar(&mysqlHost, "h", "127.0.0.1", "-h mysql host")

	flag.Parse()
}

func main() {

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDB)
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	ShowProcessList(mysqlHost, db)
	ShowInnodbLocks(mysqlHost, db)
	ShowInnodbTrx(mysqlHost, db)
	ShowInnodbLockWaits(mysqlHost, db)
}

func ShowProcessList(rds string, db *gorm.DB) {
	var pl []PROCESSLIST
	tb := "PROCESSLIST"
	dir := "/tmp/"
	fName := fmt.Sprintf("%s%s-%s.csv", dir, rds, tb)

	f, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//db.Table("PROCESSLIST").Where(`COMMAND != 'Sleep'`).Find(&pl)
	db.Table(tb).Find(&pl)

	t := table.NewWriter()
	t.SetOutputMirror(f)

	t.AppendHeader(table.Row{"ID", "USER", "HOST", "DB", "COMMAND", "TIME", "STATE", "INFO"})

	for _, j := range pl {
		t.AppendRow([]interface{}{j.ID, j.USER, j.HOST, j.DB, j.COMMAND, j.TIME, j.STATE, j.INFO})
	}

	t.RenderCSV()
}

func ShowInnodbLocks(rds string, db *gorm.DB) {
	var il []INNODB_LOCKS
	tb := "INNODB_LOCKS"
	dir := "/tmp/"
	fName := fmt.Sprintf("%s%s-%s.csv", dir, rds, tb)
	f, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	db.Table(tb).Find(&il)

	t := table.NewWriter()
	t.SetOutputMirror(f)

	t.AppendHeader(table.Row{"LockId", "LockTrxId", "LockMode", "LockType", "LockTable", "LockIndex", "LockSpace", "LockPage", "LockRec", "LockData"})
	for _, j := range il {
		t.AppendRow([]interface{}{j.LockId, j.LockTrxId, j.LockMode, j.LockType, j.LockTable, j.LockIndex, j.LockSpace, j.LockPage, j.LockRec, j.LockData})
	}
	t.RenderCSV()
}
func ShowInnodbTrx(rds string, db *gorm.DB) {
	var it []INNODB_TRX
	tb := "INNODB_TRX"
	dir := "/tmp/"
	fName := fmt.Sprintf("%s%s-%s.csv", dir, rds, tb)
	f, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	db.Table(tb).Find(&it)

	t := table.NewWriter()
	t.SetOutputMirror(f)

	t.AppendHeader(table.Row{"TrxId", "TrxState", "TrxStarted", "TrxRequestedLockId", "TrxWaitStarted", "TrxWeight", "TrxMysqlThreadId", "TrxQuery", "TrxOperationState", "TrxTablesInUse", "TrxTablesLocked", "TrxLockStructs", "TrxLockMemoryBytes", "TrxRowsLocked", "TrxRowsModified", "TrxConcurrencyTickets", "TrxIsolationLevel", "TrxUniqueChecks", "TrxForeignKeyChecks", "TrxLastForeignKeyError", "TrxAdaptiveHashLatched", "TrxAdaptiveHashTimeout", "TrxIsReadOnly", "TrxAutocommitNonLocking"})
	for _, j := range it {
		t.AppendRow([]interface{}{j.TrxId, j.TrxState, j.TrxStarted, j.TrxRequestedLockId, j.TrxWaitStarted, j.TrxWeight, j.TrxMysqlThreadId, j.TrxQuery, j.TrxOperationState, j.TrxTablesInUse, j.TrxTablesLocked, j.TrxLockStructs, j.TrxLockMemoryBytes, j.TrxRowsLocked, j.TrxRowsModified, j.TrxConcurrencyTickets, j.TrxIsolationLevel, j.TrxUniqueChecks, j.TrxForeignKeyChecks, j.TrxLastForeignKeyError, j.TrxAdaptiveHashLatched, j.TrxAdaptiveHashTimeout, j.TrxIsReadOnly, j.TrxAutocommitNonLocking})
	}
	t.RenderCSV()
}

func ShowInnodbLockWaits(rds string, db *gorm.DB) {
	var ilw []INNODB_LOCK_WAITS
	tb := "INNODB_LOCK_WAITS"
	dir := "/tmp/"
	fName := fmt.Sprintf("%s%s-%s.csv", dir, rds, tb)
	f, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	db.Table(tb).Find(&ilw)

	t := table.NewWriter()
	t.SetOutputMirror(f)

	t.AppendHeader(table.Row{"RequestingTrxId", "RequestedLockId", "BlockingTrxId", "BlockingLockId"})
	for _, j := range ilw {
		t.AppendRow([]interface{}{j.RequestingTrxId, j.RequestedLockId, j.BlockingTrxId, j.BlockingLockId})
	}
	t.RenderCSV()

}
