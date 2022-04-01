package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	cv "github.com/nirasan/go-oauth-pkce-code-verifier"
	"github.com/skratchdot/open-golang/open"
)

// AuthorizeUser implements the PKCE OAuth2 flow.
func AuthorizeUser(clientID string, authDomain string, redirectURL string) {
	// initialize the code verifier
	var CodeVerifier, _ = cv.CreateCodeVerifier()

	// Create code_challenge with S256 method
	codeChallenge := CodeVerifier.CodeChallengeS256()

	// construct the authorization URL (with Auth0 as the authorization provider)
	authorizationURL := fmt.Sprintf(
		"%s/authorize?"+
			"&response_type=code&client_id=%s"+
			"&code_challenge=%s"+
			"&code_challenge_method=S256"+
			"&redirect_uri=%s"+
			"&scope=read+write",
		authDomain, clientID, codeChallenge, redirectURL)

	// start a web server to listen on a callback URL
	server := &http.Server{Addr: ":4000"}

	// define a handler that will get the authorization code, call the token endpoint, and close the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get the authorization code
		code := r.URL.Query().Get("code")
		if code == "" {
			fmt.Println("snap: Url Param 'code' is missing")
			io.WriteString(w, "Error: could not find 'code' URL parameter\n")

			// close the HTTP server and return
			cleanup(server)
			return
		}

		// trade the authorization code and the code verifier for an access token
		codeVerifier := CodeVerifier.String()
		token, err := getAccessToken(clientID, codeVerifier, code, redirectURL)
		if err != nil {
			fmt.Println("snap: could not get access token")
			io.WriteString(w, "Error: could not retrieve access token\n")

			// close the HTTP server and return
			cleanup(server)
			return
		}

		storeToken(token)

		// return an indication of success to the caller
		io.WriteString(w, `
		<html>
			<body>
				<h1>Login successful</h1>
			</body>
		</html>`)

		cleanup(server)
	})

	// parse the redirect URL for the port number
	u, err := url.Parse(redirectURL)
	if err != nil {
		fmt.Printf("snap: bad redirect URL: %s\n", err)
		os.Exit(1)
	}

	// set up a listener on the redirect port
	port := fmt.Sprintf(":%s", u.Port())
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("snap: can't listen to port %s: %s\n", port, err)
		os.Exit(1)
	}

	// open a browser window to the authorizationURL
	err = open.Start(authorizationURL)
	if err != nil {
		fmt.Printf("snap: can't open browser to URL %s: %s\n", authorizationURL, err)
		os.Exit(1)
	}

	// start the blocking web server loop
	// this will exit when the handler gets fired and calls server.Close()
	server.Serve(l)
}

func storeToken(token string) error {
	cachePath, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	tokenCachePath := path.Join(cachePath, "phrase/token")
	return os.WriteFile(tokenCachePath, []byte(token), 0644)
}

func ReadToken() (string, error) {
	cachePath, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	tokenPath := path.Join(cachePath, "phrase/token")
	file, err := os.Open(tokenPath)
	if err != nil {
		return "", err
	}

	token, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func TokenExpired(token string) (bool, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return true, errors.New("token not valid")
	}

	return false, nil

	// payload := parts[1]

	// decoded, err := base64.StdEncoding.DecodeString(payload)
	// if err != nil {
	// 	fmt.Printf("error 1: %s\n", payload)
	// 	fmt.Println(err)
	// 	return true, err
	// }

	// var payloadData map[string]interface{}
	// err = json.Unmarshal(decoded, &payloadData)
	// if err != nil {
	// 	fmt.Println("error 2")
	// 	return true, err
	// }

	// exp := payloadData["exp"].(int64)

	// fmt.Println(exp)
	// fmt.Println(time.Now().Unix())

	// fmt.Println("valid token")
	// return time.Now().Unix() > exp, nil
}

// getAccessToken trades the authorization code retrieved from the first OAuth2 leg for an access token
func getAccessToken(clientID string, codeVerifier string, authorizationCode string, callbackURL string) (string, error) {
	// set the url and form-encoded data for the POST to the access token endpoint
	url := "http://localhost:3000/oauth/token"
	data := fmt.Sprintf(
		"grant_type=authorization_code&client_id=%s"+
			"&code_verifier=%s"+
			"&code=%s"+
			"&redirect_uri=%s",
		clientID, codeVerifier, authorizationCode, callbackURL)
	payload := strings.NewReader(data)

	// create the request and execute it
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("snap: HTTP error: %s", err)
		return "", err
	}

	// process the response
	defer res.Body.Close()
	var responseData map[string]interface{}
	body, _ := ioutil.ReadAll(res.Body)

	// unmarshal the json into a string map
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Printf("snap: JSON error: %s", err)
		return "", err
	}

	// retrieve the access token out of the map, and return to caller
	accessToken := responseData["access_token"].(string)
	return accessToken, nil
}

// cleanup closes the HTTP server
func cleanup(server *http.Server) {
	// we run this as a goroutine so that this function falls through and
	// the socket to the browser gets flushed/closed before the server goes away
	go server.Close()
}
