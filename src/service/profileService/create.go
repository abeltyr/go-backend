package profileService

import (
	"adtec/backend/src/prisma/profile/db"
	"adtec/backend/src/utils/database"
	"log"
)

func Create(input CreateInput) (*db.ProfilesModel, error) {

	client, ctx, err := database.GetProfileClient()
	if err != nil {
		return nil, err
	}
	parameter := []db.ProfilesSetParam{}

	if input.Country != nil {
		parameter = append(parameter, db.Profiles.Country.Set(*input.Country))
	}
	if input.Address != nil {
		parameter = append(parameter, db.Profiles.Address.Set(*input.Address))
	}

	if input.City != nil {
		parameter = append(parameter, db.Profiles.City.Set(*input.City))
	}

	if input.BirthDate != nil {
		parameter = append(parameter, db.Profiles.BirthDate.Set(*input.BirthDate))
	}

	createdUser, err := client.Profiles.CreateOne(
		db.Profiles.UserID.Set(input.UserId),
		db.Profiles.FullName.Set(input.FullName),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		log.Println("user creation database side error", input.FullName, "error:", err)
		return nil, err
	}

	return createdUser, nil

}
