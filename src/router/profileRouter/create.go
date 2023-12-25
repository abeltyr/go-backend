package profileRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/profileService"
	"adtec/backend/src/utils"
	"context"
	"log"
	"time"
)

func Create(ctx context.Context, input model.CreateProfileInput) (*model.Profile, error) {

	// TODO: setup userId not being register check

	var birthDate *time.Time
	if input.BirthDate != nil {
		date, _ := time.Parse("2006-01-02", *input.BirthDate)
		birthDate = &date
	}

	// TODO: setup age filter

	profileData, err := profileService.Create(profileService.CreateInput{
		UserId:    input.UserID,
		FullName:  input.FullName,
		BirthDate: birthDate,
		Country:   input.Country,
		Address:   input.Address,
		City:      input.City,
	})
	if err != nil {
		log.Println("error userService Create", err)
		return nil, err
	}

	profile := utils.ProfileGraphqlConverter(profileData)
	return profile, nil
}
