package controller

import (
	"cs-lab-6/pkg/sso/service/facebook"
	"cs-lab-6/pkg/sso/service/github"
	"cs-lab-6/pkg/sso/service/google"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	authFacebook "golang.org/x/oauth2/facebook"
	authGithub "golang.org/x/oauth2/github"
	authGoogle "golang.org/x/oauth2/google"
)

type IController interface {
	RegisterRoutes(r *gin.Engine)
}

type controller struct {
}

func New() IController {
	return &controller{}
}

func (c *controller) RegisterRoutes(r *gin.Engine) {
	r.GET("/", HandleIndex)

	googleUrl := "google"
	facebookUrl := "facebook"
	githubUrl := "github"

	googleApi := r.Group("/" + googleUrl)
	facebookApi := r.Group("/" + facebookUrl)
	githubApi := r.Group("/" + githubUrl)

	port := viper.GetString("port")
	state := viper.GetString("oauthStateString")

	googleConfig := &oauth2.Config{
		ClientID:     viper.GetString(googleUrl + ".clientID"),
		ClientSecret: viper.GetString(googleUrl + ".clientSecret"),
		RedirectURL:  "https://localhost:" + port + "/" + googleUrl + "/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     authGoogle.Endpoint,
	}
	facebookConfig := &oauth2.Config{
		ClientID:     viper.GetString(facebookUrl + ".clientID"),
		ClientSecret: viper.GetString(facebookUrl + ".clientSecret"),
		RedirectURL:  "https://localhost:" + port + "/" + facebookUrl + "/callback",
		Scopes:       []string{"public_profile"},
		Endpoint:     authFacebook.Endpoint,
	}
	githubConfig := &oauth2.Config{
		ClientID:     viper.GetString(githubUrl + ".clientID"),
		ClientSecret: viper.GetString(githubUrl + ".clientSecret"),
		RedirectURL:  "https://localhost:" + port + "/" + githubUrl + "/callback",
		Scopes:       []string{"public_profile"},
		Endpoint:     authGithub.Endpoint,
	}

	googleService := google.NewService(googleConfig, state, "/")
	facebookService := facebook.NewService(facebookConfig, state, "/")
	githubService := github.NewService(githubConfig, state, "/")

	RegisterEndpoints(googleApi, googleService)
	RegisterEndpoints(facebookApi, facebookService)
	RegisterEndpoints(githubApi, githubService)
}
