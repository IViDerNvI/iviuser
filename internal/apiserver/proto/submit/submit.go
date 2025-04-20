package submit

import (
	v1 "github.com/ividernvi/iviuser/model/v1"
)

func (s *UpdateSubmitRequest) ToSubmit() *v1.Submit {
	return &v1.Submit{
		Status:  s.Status,
		Details: s.Details,
	}
}
