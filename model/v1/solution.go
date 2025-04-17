package model

type Solution struct {
	ObjMeta `json:",inline"`

	ProblemID     string `json:"problem_id" gorm:"column:problem_id" validate:"required"`
	TestData      string `json:"data_test" gorm:"column:data_test" validate:"required"`
	TestResult    string `json:"result_test" gorm:"column:result_test" validate:"required"`
	TestTimeLimit string `json:"limit_test" gorm:"column:limit_test" validate:"required"`
}

type SolutionList struct {
	ListMeta `json:",inline"`
	Items    []Solution `json:"items"`
}
