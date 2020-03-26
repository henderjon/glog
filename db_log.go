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
// A generic table might look like this
//
// CREATE TABLE `glog_log`  (
// 	`log_id` int(11) NOT NULL AUTO_INCREMENT,
// 	`logged_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "In order to keep this table/index managable, when inserting this field, set the seconds to 00.",
// 	`location` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
// 	`message` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
// 	`context` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
// 	`created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
// 	PRIMARY KEY (`log_id`),
// 	KEY `idx_logged_at`(`logged_at`),
// 	KEY `idx_location`(`location`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
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

	sql := fmt.Sprintf(`INSERT INTO %s (logged_at, location, message, context) VALUES (?, ?, ?, ?)`, l.Table)
	_, err = l.Conn.Exec(sql, e.Timestamp.Format(GoMySQLIdxTimestamp), e.Location, e.Message, ctx)
	if err != nil {
		fmt.Println(err, e.Message)
	}
}
