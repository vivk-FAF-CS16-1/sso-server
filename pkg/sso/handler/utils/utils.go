package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HandleLogin(c *gin.Context, config *oauth2.Config, state string) {
	URL, err := url.Parse(config.Endpoint.AuthURL)
	if err != nil {
		panic("Parse: " + err.Error())
	}

	parameters := url.Values{}
	parameters.Add("client_id", config.ClientID)
	parameters.Add("scope", strings.Join(config.Scopes, " "))
	parameters.Add("redirect_uri", config.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", state)
	URL.RawQuery = parameters.Encode()
	urlString := URL.String()
	http.Redirect(c.Writer, c.Request, urlString, http.StatusTemporaryRedirect)
}

func HandleCallback(c *gin.Context, config *oauth2.Config, state string,
	redirectRout string, accessTokenUrl string, afterAccessTokenUrl string) {
	newState := c.Request.FormValue("state")

	if newState != state {
		fmt.Println("invalid oauth newState, expected " + state + ", got " + newState + "\n")
		http.Redirect(c.Writer, c.Request, redirectRout, http.StatusTemporaryRedirect)
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

	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
		return
	}

	tokenUrl := accessTokenUrl + url.QueryEscape(token.AccessToken) + afterAccessTokenUrl
	resp, err := http.Get(tokenUrl)
	if err != nil {
		http.Redirect(c.Writer, c.Request, redirectRout, http.StatusTemporaryRedirect)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll: " + err.Error() + "\n")
		http.Redirect(c.Writer, c.Request, redirectRout, http.StatusTemporaryRedirect)
		return
	}

	_, _ = fmt.Fprintf(c.Writer, "Response: %s", content)
	return

}
