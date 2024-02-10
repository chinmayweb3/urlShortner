package model

type TUrl struct {
	URL         string `json:"url,omitempty" bson:"url"`
	CustomShort string `json:"short"  bson:"customShort"`
}
