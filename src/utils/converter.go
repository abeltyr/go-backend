package utils

import (
	"adtec/backend/src/graphql/model"
	"encoding/json"
)

func LimitSetter(limit int32) int32 {
	if limit == 0 || limit >= 50 {
		limit = 50
	}
	return limit
}

func ProfileGraphqlConverter(userData interface{}) *model.Profile {
	j, _ := json.Marshal(userData)

	var user *model.Profile
	_ = json.Unmarshal(j, &user)

	return user
}
