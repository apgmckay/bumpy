package client

import (
	"bumpy/package/server/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	_ = iota
	v1
)

type Client struct {
	URL        string
	httpClient http.Client
}

func New(endpoint, timeDurationString string) (Client, error) {
	parsedEndpoint, err := url.ParseRequestURI(endpoint)
	if err != nil {
		return Client{}, err
	}

	timeout, err := time.ParseDuration(timeDurationString)
	if err != nil {
		return Client{}, err
	}

	return Client{
		URL: parsedEndpoint.String(),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c Client) BumpMajor(version string) (string, error) {
	endpoint := fmt.Sprintf("%s/api/v%d/major/%s", c.URL, v1, version)
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error")
	}

	var bumpedVersion responses.BumpedVersion

	json.Unmarshal(body, &bumpedVersion)

	return bumpedVersion.Version, nil
}

func (c Client) BumpMinor(version string) (string, error) {
	endpoint := fmt.Sprintf("%s/api/v%d/minor/%s", c.URL, v1, version)
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error")
	}

	var bumpedVersion responses.BumpedVersion

	json.Unmarshal(body, &bumpedVersion)

	return bumpedVersion.Version, nil
}

func (c Client) BumpPatch(version string) (string, error) {
	endpoint := fmt.Sprintf("%s/api/v%d/patch/%s", c.URL, v1, version)
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error")
	}

	var bumpedVersion responses.BumpedVersion

	json.Unmarshal(body, &bumpedVersion)

	return bumpedVersion.Version, nil
}
