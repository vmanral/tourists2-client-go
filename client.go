package tourists3

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"
	"strings"
)

// HostURL - Default Tourists URL
const HostURL string = "https://gorest.co.in/public/v2/users"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token    string `json:"token"`
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id`
	Username string `json:"username`
	Token    string `json:"token"`
}

// NewClient -
func NewClient(host *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Tourists URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	c.Token = "d0276fb9a596e9eafc7f786a6ce6db646bbb7915568a7adff372f3ca4268c588"

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

// GetTourists - Returns list of tourists (no auth required)
func (c *Client) GetTourists() ([]Tourists, error) {
	var tourists Tourists

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/public/v2/users", c.HostURL), nil)
	if err != nil {
		return tourists, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return tourists, err
	}

	//tourists := []Tourist{}
	err = json.Unmarshal(body, &tourists)
	if err != nil {
		return tourists, err
	}

	return tourists, nil
}

// GetTourist - Returns specific tourist (no auth required)
func (c *Client) GetTourist(touristID string) ([]Tourist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/Tourist/%s", c.HostURL, touristID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	tourists := []Tourist{}
	err = json.Unmarshal(body, &tourists)
	if err != nil {
		return nil, err
	}

	return tourists, nil
}

// CreateTourist - Create new tourist
func (c *Client) CreateTourist(tourist Tourist, authToken *string) (*Tourist, error) {
	rb, err := json.Marshal(tourist)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/Tourist", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	newTourist := Tourist{}
	err = json.Unmarshal(body, &newTourist)
	if err != nil {
		return nil, err
	}

	return &newTourist, nil
}
