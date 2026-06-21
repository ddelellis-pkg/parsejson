package parsejson
import (
	"os"
	"encoding/json"
)

// Given a path to a key:value json, returns a map of strings
func MapOfStrings(path string) (guy map[string]string, err error) {
	var data []byte
	if data, err = os.ReadFile(path); err != nil {
		return
	}
	guy = make(map[string]string)
	err = json.Unmarshal(data, &guy)
	return
}
