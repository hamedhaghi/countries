# Countries

`Countries` is a Go package for interacting with the [REST Countries API](https://restcountries.com/). It provides functions to fetch country data by name, region, capital, and language.

## Installation

Install the package using:

```go
go get github.com/hamedhaghi/countries
```

## Usage

Import the package:

```go
import "github.com/hamedhaghi/countries"
```

### Examples

Fetch all countries

```go
countries, err := countries.All()
```

Fetch by name

```go
country, err := countries.ByName("Germany")
```

Fetch by region

```go
countries, err := countries.ByRegion("Europe")
```

Fetch by capital

```go
country, err := countries.ByCapital("Berlin")
```

Fetch by language

```go
countries, err := countries.ByLanguage("German")
```

## License

This package is licensed under the MIT License.
