package reader

import (
	"fmt"
	"io"
	"log"
	"os"
)

func DoesStdinHaveData() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	if os.Getenv("GITIME_NO_STDIN") == "1" {
		return false
	}

	// Both yield false positives in CI, careful
	//if (fileInfo.Mode() & os.ModeCharDevice) == 0 { // alternatively?
	if (fileInfo.Mode() & os.ModeNamedPipe) != 0 {
		return true
	}

	return false
}

func ReadStdin() string {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", stdin)
}
