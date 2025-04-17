package model

type Submit struct {
	ObjMeta `json:",inline"`

	CodeText  string `json:"code_text" gorm:"column:code_text" validate:"required"`
	Language  string `json:"language" gorm:"column:language" validate:"required"`
	ProblemID string `json:"problem_id" gorm:"column:problem_id" validate:"required"`
	Author    string `json:"author" gorm:"column:author" validate:"required"`
}

type SubmitList struct {
	ListMeta `json:",inline"`
	Items    []Submit `json:"items"`
}

func (s *Submit) TableName() string {
	return "submits"
}
