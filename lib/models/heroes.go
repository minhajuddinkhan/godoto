package models

//Hero model
type Hero struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name"`
	LocalName   string `bson:"localized_name" json:"localized_name"`
	PrimaryAtrr string `json:"primary_attr"`
	AttackType  string `json:"attack_type"`
	Legs        int    `json:"legs"`
}
