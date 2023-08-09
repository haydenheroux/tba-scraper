module haydenheroux.github.io/tbascraper

go 1.18

replace haydenheroux.github.io/tba => ./pkg/tba

replace haydenheroux.github.io/scout => ./pkg/scout

replace haydenheroux.github.io/adapter => ./pkg/adapter

require (
	haydenheroux.github.io/adapter v0.0.0-00010101000000-000000000000
	haydenheroux.github.io/scout v0.0.0-00010101000000-000000000000
	haydenheroux.github.io/tba v0.0.0-00010101000000-000000000000
)
