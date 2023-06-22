package logger

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type DBLog struct {
	Conn  *sql.DB
	Table string
}

// NewDBLog creates a new logger that writes to the given table on the given connection
func NewDBLog(conn *sql.DB, table string) *DBLog {
	return &DBLog{
		Conn:  conn,
		Table: table,
	}
}

// Log fulfills the Logger interface. It writes the entry to the underlying destination
func (l DBLog) Log(args ...interface{}) {
	e := entry(args...)
	l.log(e)
}

// Fatal fulfills the Logger interface. It writes the entry to the underlying destination then exits
func (l DBLog) Fatal(args ...interface{}) {
	l.Log(args...)
	os.Exit(1)
}

func (l DBLog) Write(p []byte) (n int, err error) {
	e := entry(p)
	l.log(e)
	return len(p), nil
}

// log is the internal guts of the DB calls for the DBLog
//
// # A generic table might look like this
//
// CREATE TABLE "logs" (
//
//	"log_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
//	"message" TEXT,
//	"location" TEXT,
//	"created" text,
//	"level" INTEGER,
//	"context" TEXT
//
// );
func (l DBLog) log(e *Entry) {
	if e == nil || l.Conn == nil {
		return
	}

	ctx, err := json.Marshal(e.Context)
	if err != nil {
		log.Println(err)
	}

	if e.Timestamp.IsZero() {
		e.Timestamp = time.Now().UTC()
	}

	sql := fmt.Sprintf(`INSERT INTO %s (message, location, created, level, context) VALUES (?, ?, ?, ?, ?)`, l.Table)
	_, err = l.Conn.Exec(sql, e.Message, e.Location, e.Timestamp.UTC().Format(time.RFC3339), e.Level, ctx)
	if err != nil {
		fmt.Println(err, e.Message)
	}
}
