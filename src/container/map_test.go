package te

import (
	"fmt"
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[string]int{"one": 1,
		"two":   2,
		"three": 3,
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
	m1["four"] = 4
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}
