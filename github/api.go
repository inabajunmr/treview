package github

// Span is span for trending search
type Span int

const (
	// Today span
	Today Span = iota
	// Week span
	Week
	// Month span
	Month
)

// Repository is expression type for github repository on trending.
type Repository struct {
	Name        string
	Description string
	Lang        string
	Star        int
	StarBySpan  int
	Fork        int
}

func find(lang string, span Span) []Repository {
	// assemble url for trending
	// Get by http
	// parse html and struct Repositories
	return nil
}
