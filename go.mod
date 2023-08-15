module github.com/haydenheroux/tbascraper

go 1.18

replace github.com/haydenheroux/tba => ./pkg/tba

replace github.com/haydenheroux/scout => ./pkg/scout

replace github.com/haydenheroux/adapter => ./pkg/adapter

replace github.com/haydenheroux/data => ./pkg/data

require (
	github.com/haydenheroux/adapter v0.0.0-00010101000000-000000000000
	github.com/haydenheroux/data v0.0.0-00010101000000-000000000000
	github.com/haydenheroux/scout v0.0.0-00010101000000-000000000000
	github.com/haydenheroux/tba v0.0.0-00010101000000-000000000000
)
