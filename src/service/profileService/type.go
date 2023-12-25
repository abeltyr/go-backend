package profileService

import "time"

type CreateInput struct {
	UserId    string     `json:"userId"`
	FullName  string     `json:"fullName"`
	BirthDate *time.Time `json:"birthDate"`
	Country   *string    `json:"country"`
	Address   *string    `json:"address"`
	City      *string    `json:"city"`
}
