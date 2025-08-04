package sequence

import (
	"fmt"
	"io"

	"github.com/ufukty/diagramer/pkg/sequence/ast"
	"github.com/ufukty/diagramer/pkg/sequence/lexer"
)

func Render(dst io.Writer, src io.Reader) error {
	l, err := lexer.FromReader(src)
	if err != nil {
		return fmt.Errorf("lexer: %w", err)
	}

	_, err = ast.Parse(l)
	if err != nil {
		return fmt.Errorf("parsing: %w", err)
	}

	return nil
}
