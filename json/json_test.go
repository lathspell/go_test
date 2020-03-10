package json

import (
	json2 "encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Traveler struct {
	FirstName  string `json:"first_name"`
	Age        int 	  `json:"age"`
	Remarks    string `json:",omitempty"` // leading "," as "omitempty" is a flag not the fieldname!
	DatabaseId int64  `json:"-"`          // skip field in JSON
}

func TestSerialize(t *testing.T) {
	traveler := Traveler{
		FirstName:  "Tom",
		Age:        12,
		Remarks:    "",
		DatabaseId: 42421,
	}

	json,err := json2.Marshal(traveler)
	if err != nil {
		t.Errorf("Error marshalling JSON: %s", err.Error())
	}

	expected := `{"first_name":"Tom","age":12}`
	assert.Equal(t, expected, string(json))
}

func TestUnserialize(t *testing.T) {
	json := `{"first_name":"Tim","age":23}`
	traveler := Traveler{}

	if err := json2.Unmarshal([]byte(json), &traveler); err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, Traveler{FirstName: "Tim", Age: 23}, traveler)
}
