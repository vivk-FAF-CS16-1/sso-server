package facebook

import (
	"cs-lab-6/pkg/sso/handler/utils"
	"cs-lab-6/pkg/sso/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type facebookService struct {
	service.IService

	config       *oauth2.Config
	state        string
	redirectRout string
}

func NewService(config *oauth2.Config, state string, redirectRout string) service.IService {
	return &facebookService{
		config:       config,
		state:        state,
		redirectRout: redirectRout,
	}
}

func (s *facebookService) Login(c *gin.Context) {
	utils.HandleLogin(c, s.config, s.state)
}

func (s *facebookService) Callback(c *gin.Context) {
	accessTokenUrl := "https://graph.facebook.com/me?access_token="
	afterAccessTokenUrl := ""
	utils.HandleCallback(c, s.config, s.state, s.redirectRout, accessTokenUrl, afterAccessTokenUrl)
}
