//golangcitest:args -Egofmt
package testdata

import "fmt"

// want +4 "File is not properly formatted"
func _() {
	fmt.Println("foo")
}
