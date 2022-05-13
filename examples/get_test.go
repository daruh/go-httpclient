package examples

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	//arrange

	//act
	endpoints, err := GetEndpoints()
	//assert
	fmt.Println(err)
	fmt.Println(endpoints)
}
