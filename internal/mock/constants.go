package mock

import "github.com/google/uuid"

var mockStubbedUserID uuid.UUID = uuid.MustParse("0bbe5c67-d181-472e-8517-9aa491c83967")
var nonExistingUserID uuid.UUID = uuid.MustParse("4d6fa8be-2d54-4d1f-b119-60733e27ef2f")

var paymentOneID uuid.UUID = uuid.MustParse("c6c52c2f-f210-4166-9d9a-f847100879b8")
var paymentTwoID uuid.UUID = uuid.MustParse("a0742b1e-7c95-4907-89b7-7e91c3b9254c")
var paymentThreeID uuid.UUID = uuid.MustParse("691ec420-7f57-4497-86e9-dceccc615097")
var paymentFourID uuid.UUID = uuid.MustParse("96abd6bd-dde8-4d2b-951b-29268c5ff201")

var refundOneID uuid.UUID = uuid.MustParse("ff2b2d9a-2d2f-43ca-a1eb-18c2f6cbc362")
var refundTwoID uuid.UUID = uuid.MustParse("454cecd6-3eb6-4dbc-9d52-fbd5f7e69b8f")

var mockStubbedAuthToken string = "STUBBED_AUTH_TOKEN"

func MockStubbedAuthToken() string {
	return mockStubbedAuthToken
}

func MockStubbedUserID() uuid.UUID {
	return mockStubbedUserID
}

func NonExistingUserID() uuid.UUID {
	return nonExistingUserID
}

func PaymentOneID() uuid.UUID {
	return paymentOneID
}

func PaymentTwoID() uuid.UUID {
	return paymentTwoID
}

func PaymentThreeID() uuid.UUID {
	return paymentThreeID
}

func PaymentFourID() uuid.UUID {
	return paymentFourID
}

func RefundOneID() uuid.UUID {
	return refundOneID
}

func RefundTwoID() uuid.UUID {
	return refundTwoID
}
