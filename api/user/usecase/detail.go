package usecase

import (
	"github.com/bouroo/go-project-structure/api/user/repository"
	"github.com/bouroo/go-project-structure/pkg/model"
)

func ReadUserDetails(user model.UserAccount) (userDetails model.UserDetail, err error) {
	result, err := repository.ReadUserDetails(user.ID, user.Email)
	if err != nil {
		return
	}

	userDetails = model.UserDetail{
		UserAccount: model.UserAccount{
			ID:    result.ID,
			Email: result.Email,
		},
		UserProfile: model.UserProfile{
			FirstName: result.UserProfile.FirstName,
			LastName:  result.UserProfile.LastName,
			Phone:     result.UserProfile.Phone,
			Avatar:    result.UserProfile.Avatar,
		},
	}

	for _, address := range result.UserAddress {
		userDetails.UserAddress = append(userDetails.UserAddress, model.UserAddress{
			Number:   address.Number,
			Street:   address.Street,
			City:     address.City,
			Province: address.Province,
			Country:  address.Country,
			PostCode: address.PostCode,
		})
	}
	return
}
