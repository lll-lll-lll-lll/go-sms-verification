package sms_verification

type OneTimePassword struct {
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
}

type VerifyData struct {
	User *OneTimePassword `json:"user,omitempty" validate:"required"`
	Code string           `json:"code,omitempty" validate:"required"`
}
