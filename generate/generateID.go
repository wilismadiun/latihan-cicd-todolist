package generate

import "fmt"

var nextID int = 1

func GenerateID() string {
	id := fmt.Sprintf("todo-%d", nextID)
	nextID++

	return id
}
