package model

type WordModel struct {
	Word    string `xorm:"word" json:"word"`
	Meaning string `xorm:"meaning" json:"meaning"`
}

func (*WordModel) TableName() string {
	return "words"
}
