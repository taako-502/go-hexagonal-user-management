package user_secondary_adapter

import "go-hexagonal-user-management/core/domain"

func (a *userSecondaryAdapter) Delete(id int) error {
	user := &domain.User{Id: id}
	if err := a.Db.Delete(user); err != nil {
		return err.Error
	}
	return nil
}
