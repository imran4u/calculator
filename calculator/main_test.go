package calculator

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("testing Main")
	os.Exit(m.Run())
}
