package user_secondary_adapter

import "go-sample-api/application/domain"

func (a *userSecondaryAdapter) FindAll() ([]domain.User, error) {
	var users []domain.User
	result := a.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(users) == 0 {
		return nil, UserNotFoundError
	}
	return users, nil
}