package actions

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

// bodyNode returns the first <body> node in the node tree
func bodyNode(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}
	return nil, errors.New("missing <body> in the node tree")
}

// readHtmlBodyToString reads the HTML body from the reader and returns it as a string
func readHtmlBodyToString(r io.Reader) (string, error) {
	resHtml, err := html.ParseWithOptions(r, html.ParseOptionEnableScripting(false))
	if err != nil {
		return "", err
	}
	resBody, err := bodyNode(resHtml)
	if err != nil {
		return "", err
	}
	var resBodyBuf strings.Builder
	err = html.Render(&resBodyBuf, resBody)
	if err != nil {
		return "", err
	}
	return resBodyBuf.String(), nil
}

func (as *ActionSuite) Test_Auth_LogInSignupPage() {
	res := as.HTML("/auth").Get()
	root := as.HTML("/").Get()

	as.Equal(http.StatusOK, res.Code)

	resHtmlBody, err := readHtmlBodyToString(res.Body)
	as.NoError(err)
	rootHtmlBody, err := readHtmlBodyToString(root.Body)
	as.NoError(err)

	// HTML content of auth page should be the same as root page because this is an SPA
	as.Equal(rootHtmlBody, resHtmlBody)
}

func (as *ActionSuite) Test_Auth_SignUp_Success() {
	// 1. Log in and get a token.

	as.LoadFixture("user_basic")

	res := as.JSON("/auth/signup").Post(map[string]string{
		"Username": "NeWUser",
		"Email":    "brand-new@example.com",
		"Password": "azerty",
	})

	as.Equal(http.StatusCreated, res.Code)
	var response map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&response)

	as.NoError(err)

	token, hasToken := response["token"]
	as.True(hasToken, "response should have a 'token' field")
	as.NotEmpty(token)
	as.Equal("newuser", response["username"])
	as.Equal("brand-new@example.com", response["email"])
	as.Len(response, 3, "extra fields in response")

	// 2. Validate the token by restoring a session from it.

	res = as.JSON("/auth/restore-session").Post(map[string]string{
		"Token": token.(string),
	})

	as.Equal(http.StatusOK, res.Code)
	var validateResponse map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&validateResponse)

	as.NoError(err)

	token, hasToken = validateResponse["token"]
	as.True(hasToken, "validateResponse should have a 'token' field")
	as.NotEmpty(token)
	as.Equal("newuser", validateResponse["username"])
	as.Equal("brand-new@example.com", validateResponse["email"])
	as.Len(validateResponse, 3, "extra fields in response")
}

func (as *ActionSuite) Test_Auth_LogIn_Success() {
	// 1. Log in and get a token.

	as.LoadFixture("user_basic")

	u, err := as.Users.FindByNanoID("user-5001")
	as.NoError(err)

	res := as.JSON("/auth/login").Post(map[string]string{
		"Email":    u.Email,
		"Password": "12345678",
	})

	as.Equal(http.StatusOK, res.Code)
	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)

	as.NoError(err)

	token, hasToken := response["token"]
	as.True(hasToken, "response should have a 'token' field")
	as.NotEmpty(token)
	as.Equal(u.Username, response["username"])
	as.Equal(u.Email, response["email"])
	as.Len(response, 3, "extra fields in response")

	// 2. Validate the token by restoring a session from it.

	res = as.JSON("/auth/restore-session").Post(map[string]string{
		"Token": token.(string),
	})

	as.Equal(http.StatusOK, res.Code)
	var validateResponse map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&validateResponse)

	as.NoError(err)

	token, hasToken = validateResponse["token"]
	as.True(hasToken, "validateResponse should have a 'token' field")
	as.NotEmpty(token)
	as.Equal(u.Username, validateResponse["username"])
	as.Equal(u.Email, validateResponse["email"])
	as.Len(validateResponse, 3, "extra fields in response")
}

func (as *ActionSuite) Test_Auth_LogIn_InvalidEmail() {
	as.LoadFixture("user_basic")

	res := as.JSON("/auth/login").Post(map[string]string{
		"Email":    "thisisnotanemail@example.com",
		"Password": "12345678",
	})

	as.Equal(http.StatusUnauthorized, res.Code)
	var response map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&response)

	as.NoError(err)
	as.Equal(map[string]any{
		"errors": map[string]any{
			"credentials": []any{"invalid email or password"},
		},
	}, response)
}

func (as *ActionSuite) Test_Auth_LogIn_InvalidPassword() {
	as.LoadFixture("user_basic")

	u, err := as.Users.FindByNanoID("user-5001")
	as.NoError(err)

	res := as.JSON("/auth/login").Post(map[string]string{
		"Email":    u.Email,
		"Password": "thisisnotthepassword",
	})

	as.Equal(http.StatusUnauthorized, res.Code)
	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)

	as.NoError(err)
	as.Equal(map[string]any{
		"errors": map[string]any{
			"credentials": []any{"invalid email or password"},
		},
	}, response)
}
