package usphonenumbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	// Valid phone numbers with Country Code
	phone, err := Parse("14155552233")
	assert.NotNil(t, phone)
	assert.Equal(t, "415", phone.AreaCode)
	assert.Equal(t, "4155552233", phone.PhoneNumber)
	assert.NoError(t, err)
	// Valid phone numbers with no country code
	phone, err = Parse("4155552233")
	assert.NotNil(t, phone)
	assert.Equal(t, "415", phone.AreaCode)
	assert.Equal(t, "4155552233", phone.PhoneNumber)
	assert.NoError(t, err)
	// Valid phone numbers with additional non-numeric characters
	phone, err = Parse("(415)555-2233")
	assert.NotNil(t, phone)
	assert.Equal(t, "415", phone.AreaCode)
	assert.Equal(t, "4155552233", phone.PhoneNumber)
	assert.NoError(t, err)
	phone, err = Parse("(415) 555 - 2233 A")
	assert.NotNil(t, phone)
	assert.Equal(t, "415", phone.AreaCode)
	assert.Equal(t, "4155552233", phone.PhoneNumber)
	assert.NoError(t, err)
	// Invalid length
	phone, err = Parse("155552233")
	assert.Nil(t, phone)
	assert.Error(t, err)
	// Invalid country code
	phone, err = Parse("24155552233")
	assert.Nil(t, phone)
	assert.Error(t, err)
	// Invalid area code
	phone, err = Parse("0155552233")
	assert.Nil(t, phone)
	assert.Error(t, err)
}
