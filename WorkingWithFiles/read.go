package WorkingWithFiles

import (
	"fmt"
)

func Read(name string) string {
	return fmt.Sprintf("Hello %s", name)
}