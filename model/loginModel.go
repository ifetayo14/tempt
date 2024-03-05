package model

type (
	Request struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LogoutRequest struct {
		AccessToken string `json:"accessToken"`
	}
)
