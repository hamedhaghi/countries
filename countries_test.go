package countries_test

import (
	"strings"
	"testing"

	"github.com/hamedhaghi/countries"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	countries, err := countries.All()
	assert.NoError(t, err)
	assert.NotEmpty(t, countries)
}

func TestName(t *testing.T) {
	country, err := countries.ByName("germany")
	assert.NoError(t, err)
	assert.Equal(t, "germany", strings.ToLower(country.Name.Common))
}

func TestRegion(t *testing.T) {
	countries, err := countries.ByRegion("europe")
	assert.NoError(t, err)
	assert.NotEmpty(t, countries)
}
