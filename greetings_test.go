package chiefuslecture00

import (
	"fmt"
	"testing"
)

var messageText = "Hello 🗺️ !"

func TestGetMessage(t *testing.T) {
	message := GetMessage()
	if messageText != message {
		t.Error("GetMessage returned unexpected value")
	}
	fmt.Println(message)
}
