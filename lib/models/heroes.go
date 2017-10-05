package models

type Hero struct {
	name        string   `bson:"name" json:"name"`
	localName   string   `bson:"localized_name" json:"localized_name"`
	primaryAtrr string   `bson:"primary_attr" json:"localized_name"`
	attackType  string   `bson:"attack_type" json:"attack_type"`
	roles       []string `bson: "roles" json:"roles"`
	legs        int      `bson:"legs" json:"legs"`
}
