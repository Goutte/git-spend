package reader

import (
	"fmt"
	"io"
	"os"
)

func ReadStdin() string {
	stdin, _ := io.ReadAll(os.Stdin)
	return fmt.Sprintf("%s", stdin)
}
