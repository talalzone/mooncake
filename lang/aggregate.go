package lang

type Error struct {
	Code string
	Info string
}

type ValidationResult struct {
	Fatal   []Error `json:"fatalCodes"`
	Severe  []Error `json:"severeCodes"`
	Warning []Error `json:"warningCodes"`
}
