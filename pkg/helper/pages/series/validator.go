package series

type Validator interface {
	Name() string
	Validate(match MatchDetails, previous ...MatchDetails) map[string]error
}
