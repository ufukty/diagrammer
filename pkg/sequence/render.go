package sequence

import (
	"fmt"
	"io"

	"github.com/ufukty/diagramer/pkg/sequence/ast"
)

func Render(dst io.Writer, src io.Reader) error {
	_, err := ast.Parse(src)
	if err != nil {
		return fmt.Errorf("parsing: %w", err)
	}

	return nil
}
