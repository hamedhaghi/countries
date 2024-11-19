package countries_test

import (
	"strings"
	"testing"

	"github.com/hamedhaghi/countries"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Parallel()
	countries, err := countries.All()
	assert.NoError(t, err)
	assert.NotEmpty(t, countries)
}

func TestName(t *testing.T) {
	t.Parallel()
	country, err := countries.ByName("germany")
	assert.NoError(t, err)
	assert.Equal(t, "germany", strings.ToLower(string(country.Name.Common)))
}

func TestRegion(t *testing.T) {
	t.Parallel()
	countries, err := countries.ByRegion("europe")
	assert.NoError(t, err)
	assert.NotEmpty(t, countries)
}

func TestCapital(t *testing.T) {
	t.Parallel()
	country, err := countries.ByCapital("berlin")
	assert.NoError(t, err)
	assert.Equal(t, "berlin", strings.ToLower(country.Capital[0]))
}

func TestLanguage(t *testing.T) {
	t.Parallel()
	countries, err := countries.ByLanguage("german")
	assert.NoError(t, err)
	assert.NotEmpty(t, countries)
	assert.Equal(t, "german", strings.ToLower(countries[0].Languages["deu"]))
}

func TestFlagExists(t *testing.T) {
	t.Parallel()
	country, err := countries.ByName("germany")
	assert.NoError(t, err)
	assert.NotEmpty(t, country)
	assert.NotEmpty(t, country.Flags["png"])
	assert.NotEmpty(t, country.Flags["svg"])

}

func TestCurrencyExists(t *testing.T) {
	t.Parallel()
	country, err := countries.ByName("germany")
	assert.NoError(t, err)
	assert.NotEmpty(t, country)
	assert.NotEmpty(t, country.Currencies["EUR"])
	assert.NotEmpty(t, country.Currencies["EUR"].Symbol)
}
