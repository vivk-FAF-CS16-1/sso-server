package github

import (
	"cs-lab-6/pkg/sso/handler/utils"
	"cs-lab-6/pkg/sso/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type githubService struct {
	service.IService

	config       *oauth2.Config
	state        string
	redirectRout string
}

func NewService(config *oauth2.Config, state string, redirectRout string) service.IService {
	return &githubService{
		config:       config,
		state:        state,
		redirectRout: redirectRout,
	}
}

func (s *githubService) Login(c *gin.Context) {
	utils.HandleLogin(c, s.config, s.state)
}

func (s *githubService) Callback(c *gin.Context) {
	accessTokenUrl := "https://api.github.com/user"
	newState := c.Request.FormValue("state")

	if newState != s.state {
		fmt.Println("invalid oauth newState, expected " + s.state + ", got " + newState + "\n")
		http.Redirect(c.Writer, c.Request, s.redirectRout, http.StatusTemporaryRedirect)
		return
	}

	code := c.Request.FormValue("code")

	if code == "" {
		_, _ = c.Writer.Write([]byte("Code Not Found to provide AccessToken..\n"))
		reason := c.Request.FormValue("error_reason")
		if reason == "user_denied" {
			_, _ = c.Writer.Write([]byte("User has denied Permission.."))
		}
		return
	}

	token, err := s.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
		return
	}

	tokenUrl := accessTokenUrl // + url.QueryEscape(token.AccessToken)

	client := http.Client{}
	request, err := http.NewRequest("GET", tokenUrl, nil)
	if err != nil {
		fmt.Println("http.NewRequest() failed with " + err.Error() + "\n")
		return
	}

	request.Header.Add("Authorization", "Token "+url.QueryEscape(token.AccessToken))
	resp, err := client.Do(request)
	// resp, err := http.Get(tokenUrl)
	if err != nil {
		http.Redirect(c.Writer, c.Request, s.redirectRout, http.StatusTemporaryRedirect)
		return
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll: " + err.Error() + "\n")
		http.Redirect(c.Writer, c.Request, s.redirectRout, http.StatusTemporaryRedirect)
		return
	}

	_, _ = fmt.Fprintf(c.Writer, "Response: %s", content)
	return
}
