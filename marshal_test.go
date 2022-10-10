package set

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestMarshalYAML(t *testing.T) {
	set := NewValues(1, 2, 3)
	expected := "- 1\n- 2\n- 3\n"

	data, err := yaml.Marshal(set)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, string(data), "marshalled data not equal")

	// unmarshal
	{
		var x Set[int]
		err := yaml.Unmarshal([]byte(expected), &x)
		if err != nil {
			t.Fatal(err)
		}

		if !x.Equals(set) {
			t.Errorf("unmarshalled %s, not equal to %s", x.String(), set.String())
		}
	}
}

func TestMarshalYAMLEmpty(t *testing.T) {
	set := NewValues[int]()
	expected := "[]\n"

	data, err := yaml.Marshal(set)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, string(data), "marshalled data not equal")

	// unmarshal
	{
		var x Set[int]
		err := yaml.Unmarshal([]byte(expected), &x)
		if err != nil {
			t.Fatal(err)
		}

		if !x.Equals(set) {
			t.Errorf("unmarshalled %s, not equal to %s", x.String(), set.String())
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	set := NewValues(1, 2, 3)
	expected := "[1,2,3]"

	data, err := json.Marshal(set)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, string(data), "marshalled data not equal")
}

func TestMarshalJSONEmpty(t *testing.T) {
	set := NewValues[int]()
	expected := "[]"

	data, err := json.Marshal(set)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, string(data), "marshalled data not equal")
}
