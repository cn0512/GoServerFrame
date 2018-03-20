package MysqlStd

import (
	"database/sql"
	"database/sql/driver"
	"sync"
	"sync/atomic"

	mysql "github.com/go-sql-driver/mysql"
)

func init() {
	sql.Register("mysql_hack", &MySQLDriver{})
}

type MySQLDriver struct {
}

func (d MySQLDriver) Open(dsn string) (driver.Conn, error) {
	var real mysql.MySQLDriver
	conn, err := real.Open(dsn)
	if err != nil {
		return nil, err
	}
	return &mysqlConn{Conn: conn, stmtCache: make(map[string]*mysqlStmt)}, nil
}

type mysqlConn struct {
	driver.Conn
	stmtMutex sync.RWMutex
	stmtCache map[string]*mysqlStmt
}

func (mc *mysqlConn) Prepare(query string) (driver.Stmt, error) {
	mc.stmtMutex.RLock()
	if stmt, exists := mc.stmtCache[query]; exists {
		// must update reference counter in lock scope
		atomic.AddInt32(&stmt.ref, 1)
		mc.stmtMutex.RUnlock()
		return stmt, nil
	}
	mc.stmtMutex.RUnlock()

	mc.stmtMutex.Lock()
	defer mc.stmtMutex.Unlock()

	// double check
	if stmt, exists := mc.stmtCache[query]; exists {
		atomic.AddInt32(&stmt.ref, 1)
		return stmt, nil
	}

	stmt, err := mc.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}

	stmt2 := &mysqlStmt{stmt, mc, query, 1}
	mc.stmtCache[query] = stmt2
	return stmt2, nil
}

func (mc *mysqlConn) Begin() (driver.Tx, error) {
	tx, err := mc.Conn.Begin()
	if err != nil {
		return nil, err
	}
	return &mysqlTx{mc, tx}, nil
}

type mysqlTx struct {
	*mysqlConn
	tx driver.Tx
}

func (tx *mysqlTx) Commit() (err error) {
	return tx.tx.Commit()
}

func (tx *mysqlTx) Rollback() (err error) {
	return tx.tx.Rollback()
}

type mysqlStmt struct {
	driver.Stmt
	conn  *mysqlConn
	query string
	ref   int32
}

func (stmt *mysqlStmt) Close() error {
	stmt.conn.stmtMutex.Lock()
	defer stmt.conn.stmtMutex.Unlock()

	if atomic.AddInt32(&stmt.ref, -1) == 0 {
		delete(stmt.conn.stmtCache, stmt.query)
		return stmt.Stmt.Close()
	}
	return nil
}
