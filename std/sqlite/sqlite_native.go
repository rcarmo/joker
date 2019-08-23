package sqlite

import (
    "time"
    "github.com/bvinc/go-sqlite-lite/sqlite3"
        
	//. "github.com/candid82/joker/core"
)

func setTimeout(conn *sqlite3.Conn) error {
    conn.BusyTimeout(5 * time.Second)
    return nil;
}

