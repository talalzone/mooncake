package lang

type Error struct {
	Code string
	Info string
}

type ValidationResult struct {
	Fatal   []Error `json:"Fatal"`
	Severe  []Error `json:"Severe"`
	Warning []Error `json:"Warning"`
}

func NewValidationResult() ValidationResult {
	return ValidationResult{}
}
