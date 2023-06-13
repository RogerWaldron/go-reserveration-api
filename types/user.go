package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 15
	minLengthFirstName = 2
	minLengthLastName = 2
	minLengthPassword = 8
)

type CreateUserParams struct {
	FirstName 				string 	`json:"firstName"`
	LastName 					string	`json:"lastName"`
	Email 						string  `json:"email"`
	Password 					string  `json:"password"`
}

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minLengthFirstName {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minLengthFirstName)
	}
	if len(params.LastName) < minLengthLastName {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", minLengthLastName)
	}
	if len(params.Password) < minLengthFirstName {
		errors["password"] = fmt.Sprintf("Password length should be at least %d characters", minLengthPassword)
	}
	if !isValidEmail(params.Email) {
		errors["email"] = fmt.Sprintf("Email %s not valid", params.Email)
	}

	return errors
}

func isValidEmail(e string) bool {
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		return emailRegex.MatchString(e)
}

type User struct {
	ID 								primitive.ObjectID 	`bson:"_id,omitempty" json:"id,omitempty"`
	FirstName 				string 	`bson:"firstName" json:"firstName"`
	LastName 					string	`bson:"lastName" json:"lastName"`
	Email 						string 	`bson:"email" json:"email"`
	EncryptedPassword string	`bson:"EncryptedPassword" json:"-"`
} 

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encryptedPwd, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName: params.FirstName,
		LastName: params.LastName,
		Email: params.Email,
		EncryptedPassword: string(encryptedPwd),
	}, nil
}