package database

import (
	profileDB "adtec/backend/src/prisma/profile/db"
	"adtec/backend/src/utils/database/clients"
	"context"
)

var profileClient *profileDB.PrismaClient
var profileContext context.Context
var profileErr error

func ConnectToDatabase() {
	profileClient, profileContext, profileErr = clients.ProfilePrismaClient()
}

func GetProfileClient() (*profileDB.PrismaClient, context.Context, error) {
	if profileErr != nil || profileClient == nil || profileContext == nil {
		if profileClient != nil {
			profileClient.Prisma.Disconnect()
		}
		profileClient, profileContext, profileErr = clients.ProfilePrismaClient()
	}

	return profileClient, profileContext, profileErr
}
