package ascii

import (
	"fmt"
	"os"
	"strings"
)

func RemoveSpecialCases(s string) string {
	cases := map[string]bool{"\\a": true, "\\t": true, "\\b": true, "\\v": true, "\\r": true, "\\f": true}
	for char := range cases {
		if strings.Contains(s, char) {
			fmt.Printf("Speacial case /%q/ is not support\n", char)
			os.Exit(1)
		}
	}
	return s
}
