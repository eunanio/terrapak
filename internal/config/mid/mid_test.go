package mid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMidValidation(t *testing.T){
	mid, _ := NewMID("test", "test", "test", "0.0.1")
	err := validate(mid)
	if err != nil {
		t.Errorf("Error parsing mid: %s", err)
	}
}

func TestMidValidationEmptyName(t *testing.T){
	mid, _ := NewMID("test", "", "test", "0.0.1")
	err := validate(mid); if err != nil {
		assert.Equal(t, "[mid parsing error]: name is empty", err.Error())
		return
	}

	assert.Nil(t, err)
}

func TestMidValidationEmptyProvider(t *testing.T){
	mid, _ := NewMID("test", "test", "", "0.0.1")
	err := validate(mid); if err != nil {
		assert.Equal(t, "[mid parsing error]: provider is empty", err.Error())
		return
	}
	t.Fail()
}

func TestMidValidationEmptyVersion(t *testing.T){
	mid, _ := NewMID("test", "test", "test", "")
	_,err := buildVersion(mid.Version); if err != nil {
		return
	}
	assert.NotNil(t, err)
}

func TestMidValidationInvalidVersion(t *testing.T){
	mid, _ := NewMID("test", "test", "test", ".1")
	_,err := buildVersion(mid.Version);

	assert.Equal(t, "Invalid Semantic Version", err.Error())
}

func TestMidValidationInvalidName(t *testing.T){
	mid, _ := NewMID("test", "SELECT * FROM modules", "test", "0.0.1")
	err := validate(mid); if err != nil {
		assert.Equal(t, "invalid characters in name", err.Error())
		return
	}
	t.Fail()
}

