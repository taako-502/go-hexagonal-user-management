package user_secondary_adapter

import user_model "go-hexagonal-user-management/core/models"

func (a *userSecondaryAdapter) Delete(id int) error {
	user := &user_model.User{Id: id}
	if err := a.Db.Delete(user); err != nil {
		return err.Error
	}
	return nil
}
