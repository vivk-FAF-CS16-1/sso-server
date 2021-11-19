package google

import (
	"cs-lab-6/pkg/sso/handler/utils"
	"cs-lab-6/pkg/sso/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type googleService struct {
	service.IService

	config       *oauth2.Config
	state        string
	redirectRout string
}

func NewService(config *oauth2.Config, state string, redirectRout string) service.IService {
	return &googleService{
		config:       config,
		state:        state,
		redirectRout: redirectRout,
	}
}

func (s *googleService) Login(c *gin.Context) {
	utils.HandleLogin(c, s.config, s.state)
}

func (s *googleService) Callback(c *gin.Context) {
	accessTokenUrl := "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	utils.HandleCallback(c, s.config, s.state, s.redirectRout, accessTokenUrl, "")
}
