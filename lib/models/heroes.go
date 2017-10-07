package models

type Hero struct {
	Name        string `json:"name"`
	LocalName   string `json:"localized_name"`
	PrimaryAtrr string `json:"primary_attr"`
	AttackType  string `json:"attack_type"`
	Legs        int    `json:"legs"`
}
