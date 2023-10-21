package user_secondary_adapter

import "go-sample-api/application/domain"

func (a *userSecondaryAdapter)Delete(id int) error {
	user := &domain.User{Id: id}
	if err := a.Db.Delete(user); err != nil {
			return err.Error
	}
	return nil
}