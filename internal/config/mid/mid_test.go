package mid

import "testing"

func TestMidValidation(t *testing.T){
	mid := NewMID("test", "test", "test", "0.0.1")
	err := validate(mid)
	if err != nil {
		t.Errorf("Error parsing mid: %s", err)
	}
}