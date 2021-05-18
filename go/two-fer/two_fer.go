package twofer

import (
	"fmt"
	"strings"
)

// ShareWith formats the requested Two-fer string and 
// returns it to the caller
func ShareWith(name string) string {
	if strings.TrimSpace(name) == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
