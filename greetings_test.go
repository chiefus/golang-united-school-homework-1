package chiefuslecture00

import (
	"fmt"
	"testing"
)

var greetingsText = "Hello 🗺️ !"

func TestGreetingFunc(t *testing.T) {
	greetings := GetGreetings()
	if greetingsText != greetings {
		t.Error("GetGreetings returned unexpected value")
	}
	fmt.Println(greetings)
}
