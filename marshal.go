package set

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// Marshal to JSON array. The keys will be sorted
func (set Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.ToSortedSlice(func(x, y T) int {
		return strings.Compare(fmt.Sprint(x), fmt.Sprint(y))
	}))
}

// Unmarshal from JSON array
func (set *Set[T]) UnmarshalJSON(data []byte) error {
	var keys []T
	err := json.Unmarshal(data, &keys)
	if err != nil {
		return err
	}
	s := NewValues(keys...)
	*set = s
	return nil
}

// Marshal to YAML array. The keys will be sorted
// https://pkg.go.dev/gopkg.in/yaml.v3 MarshalYAML interface
func (set Set[T]) MarshalYAML() (interface{}, error) {
	return set.ToSortedSlice(func(x, y T) int {
		return strings.Compare(fmt.Sprint(x), fmt.Sprint(y))
	}), nil
}

// Unmarshal from YAML array
// https://pkg.go.dev/gopkg.in/yaml.v3 UnmarshalYAML interface
func (set *Set[T]) UnmarshalYAML(value *yaml.Node) error {
	var keys []T
	err := value.Decode(&keys)
	if err != nil {
		return err
	}
	s := NewValues(keys...)
	*set = s
	return nil
}
