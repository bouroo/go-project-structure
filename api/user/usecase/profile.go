package usecase

import (
	"github.com/bouroo/go-project-structure/api/user/repository"
	"github.com/bouroo/go-project-structure/pkg/entity"
	"github.com/bouroo/go-project-structure/pkg/model"
)

func UpdateUserProfile(profileID, userID string, profile model.UserProfile) (err error) {
	profileEntity := entity.UserProfile{
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Phone:     profile.Phone,
		Avatar:    profile.Avatar,
	}
	err = repository.UpdateUserProfile(profileID, userID, profileEntity)
	return
}
