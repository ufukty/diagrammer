package ast

type Stmt interface {
	_stmt()
}

func (*Activate) _stmt()     {}
func (*Alternative) _stmt()  {}
func (*Break) _stmt()        {}
func (*Create) _stmt()       {}
func (*Critical) _stmt()     {}
func (*Deactivate) _stmt()   {}
func (*Destroy) _stmt()      {}
func (*LifelineDecl) _stmt() {}
func (*Loop) _stmt()         {}
func (*Message) _stmt()      {}
func (*Note) _stmt()         {}
func (*Optional) _stmt()     {}
func (*Parallel) _stmt()     {}
func (*WideNote) _stmt()     {}
