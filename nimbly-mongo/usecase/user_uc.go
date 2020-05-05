package usecase

import (
	"chi-rest/model"
)

// GetUser ...
func (uc UC) GetUser() ([]*model.User, error) {
	dt, err := model.UserOp.Get(uc.DB)
	return dt, err
}
