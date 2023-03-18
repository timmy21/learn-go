package sql

type StmtNode interface {
	statement()
}

type SelectStmt struct {
	Table   string
	Columns []string
}

func (stmt *SelectStmt) statement() {}
