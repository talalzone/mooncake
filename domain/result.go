package domain

type Error struct {
	Code int
	Info string
}

type ValidationResult struct {
	FatalCodes   []Error `json:"fatalCodes"`
	SevereCodes  []Error `json:"severeCodes"`
	WarningCodes []Error `json:"warningCodes"`
}
