package services

import (
	"microservices-ex-app/oauth-api/src/api/domain/oauth"
	"microservices-ex-app/src/api/utils/errors"
	"time"
)

type oauthService struct {
}
type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError)
}

var OauthService oauthServiceInterface

func init() {
	OauthService = &oauthService{}
}

func (o *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	//Generate and Store the access token in a persistent way because if not we will be generating one everytime
	token := oauth.AccessToken{
		UserId:  user.Id,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}
	if err := token.Save(); err != nil {
		return nil, err
	}
	return &token, nil
}

func (o *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)
	if err != nil {
		return nil, err
	}

	if token.IsExpired() {
		return nil, errors.NewNotFoundError("no access token found for given information")
	}
	return token, err
}
