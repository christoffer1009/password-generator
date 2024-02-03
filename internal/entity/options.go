package entity

type Options struct {
	Length       int  `json:"length"`
	UpperCase    bool `json:"upperCase"`
	LowerCase    bool `json:"lowerCase"`
	Digits       bool `json:"digits"`
	SpecialChars bool `json:"specialChars"`
}
