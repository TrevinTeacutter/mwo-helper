package series

type Validator func(match MatchDetails, previous ...MatchDetails) map[string]error
