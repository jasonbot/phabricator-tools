package phabricatortools

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"path"
)

type arcrc struct {
	Config map[string](string)            `json:"config"`
	Hosts  map[string](map[string]string) `json:"hosts"`
}

func readarcrc() (arcrc, error) {
	user, _ := user.Current()
	arcrcpath := path.Join(user.HomeDir, ".arcrc")
	arcrcHandle, err := os.Open(arcrcpath)

	if err != nil {
		return arcrc{}, err
	}

	defer arcrcHandle.Close()

	bytes, err := ioutil.ReadAll(arcrcHandle)

	if err != nil {
		return arcrc{}, err
	}

	var arcrcData arcrc
	err = json.Unmarshal(bytes, &arcrcData)

	if err != nil {
		return arcrc{}, err
	}

	return arcrcData, nil
}

// GetArcEndpointAndToken will return the user's cached API endpoint/token used with arcanist
func GetArcEndpointAndToken() (string, string, error) {
	rcdata, err := readarcrc()

	if err != nil {
		return "", "", err
	}

	defaultURL, _ := url.Parse(rcdata.Config["default"])

	for key, tokenHash := range rcdata.Hosts {
		keyURL, _ := url.Parse(key)

		if defaultURL.Scheme == keyURL.Scheme && defaultURL.Host == keyURL.Host {
			return key, tokenHash["token"], nil
		}
	}

	return "", "", errors.New("Couldn't find default in .arcrc")
}

/*
// Unfortunately this doesn't work at work, the phabricator instance always redirects to the Google SSO page
func dial() (*gonduit.Conn, error) {
	endpoint, token, err := GetArcEndpointAndToken()

	if err != nil {
		return nil, err
	}

	fmt.Printf("%v \n", token)

	return gonduit.Dial(
		endpoint,
		&core.ClientOptions{
			APIToken: token,
		},
	)
}
*/
