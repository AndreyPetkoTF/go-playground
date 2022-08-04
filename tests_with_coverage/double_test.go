package double

import (
	"fmt"
	"testing"
)

/**
To open html with coverage info:
	go test -short -count=1 -race -coverprofile=coverage.out
	go tool cover -html=coverage.out

*/

func TestDouble(t *testing.T) {
	res := double(10)

	if res != 20 {
		fmt.Errorf("error: expected value: %v", 20)
	}
}
