package user_secondary_adapter

import user_model "go-hexagonal-user-management/core/models"

func (a *userSecondaryAdapter) FindAll() ([]user_model.User, error) {
	var users []user_model.User
	result := a.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(users) == 0 {
		return nil, ErrUserNotFound
	}
	return users, nil
}
