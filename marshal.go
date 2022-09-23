package set

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func (set *Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.ToSlice())
}

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

func (set *Set[T]) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(set.ToSlice())
}

func (set *Set[T]) UnmarshalYAML(data []byte) error {
	var keys []T
	err := yaml.Unmarshal(data, &keys)
	if err != nil {
		return err
	}
	s := NewValues(keys...)
	*set = s
	return nil
}
