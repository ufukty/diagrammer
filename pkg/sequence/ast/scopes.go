package ast

type ScopeDefining interface {
	AppendStmt(Stmt)
}

func (c *Case) AppendStmt(stmt Stmt)    { c.Stmts = append(c.Stmts, stmt) }
func (d *Diagram) AppendStmt(stmt Stmt) { d.Stmts = append(d.Stmts, stmt) }
