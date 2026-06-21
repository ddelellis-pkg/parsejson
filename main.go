package parsejson
import (
	"fmt"
	"os"
	"encoding/json"
)

// Given a path to a key:value json, returns a map of strings
func MapOfStrings(path string) (guy map[string]string, err error) {
	var data []byte
	if data, err = os.ReadFile(path); err != nil {
		err = fmt.Errorf("Failed parsing %s: %w", path, err)
		return
	}
	guy = make(map[string]string)
	if err = json.Unmarshal(data, &guy); err != nil {
		err = fmt.Errorf("Failed marshalling data to struct %T: %w", guy, err)
	}
	return
}
