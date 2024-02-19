package repository

import "github.com/bouroo/go-clean-arch/entity"

func (r *userRepository) CreateUserAddress(address *entity.UserAddress) (err error) {
	dbTx := r.DB.Begin()
	defer dbTx.Rollback()

	err = dbTx.Create(address).Error
	if err != nil {
		r.Logger.Error("CreateUserAddress", "error", err)
		return err
	}

	return dbTx.Commit().Error
}

func (r *userRepository) ReadUserAddress(userID, addressID string) (addresses []entity.UserAddress, err error) {

	err = r.DB.Where(&entity.UserAddress{ID: addressID, UserID: userID}).Find(&addresses).Error
	return
}

func (r *userRepository) UpdateUserAddress(addressID string, address entity.UserAddress) (err error) {

	dbTx := r.DB.Begin()
	defer dbTx.Rollback()

	err = dbTx.Model(&entity.UserAddress{ID: addressID}).Updates(address).Error
	if err != nil {
		r.Logger.Error("UpdateUserAddress", "error", err)
		return err
	}

	return dbTx.Commit().Error
}

func (r *userRepository) DeleteUserAddress(addressID string) (err error) {

	dbTx := r.DB.Begin()
	defer dbTx.Rollback()

	err = dbTx.Delete(&entity.UserAddress{}, addressID).Error
	if err != nil {
		r.Logger.Error("DeleteUserAddress", "error", err)
		return err
	}

	return dbTx.Commit().Error
}
