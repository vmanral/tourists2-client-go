package tourists

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetAllOrders - Returns all user's order
func (c *Client) GetAllOrders(authToken *string) (*[]Order, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/orders", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	orders := []Order{}
	err = json.Unmarshal(body, &orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

// GetOrder - Returns a specifc order
func (c *Client) GetOrder(orderID string, authToken *string) (*Order, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/Tourist/%s", c.HostURL, orderID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	order := Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// GetSpecificTourist - Returns details of a specifc tourist
func (c *Client) GetSpecificTourist(touristID string) (*TouristData, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/Tourist/%s", c.HostURL, touristID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	order := TouristData{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// CreateOrder - Create new order
func (c *Client) CreateOrder(orderItems []OrderItem, authToken *string) (*Order, error) {
	rb, err := json.Marshal(orderItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/Tourist", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	order := Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// CreateNewTourist - Create new tourist
func (c *Client) CreateNewTourist(orderItems TouristInput) (*TouristData, error) {
	fmt.Println(orderItems)
	rb, err := json.Marshal(orderItems)
	fmt.Println(rb)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(rb))
	fmt.Println(strings.NewReader(string(rb)))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/Tourist", c.HostURL), strings.NewReader(string(rb)))
	req.Header.Add("Content-Type", "application/json")
	fmt.Println(req)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	fmt.Println(body)
	if err != nil {
		return nil, err
	}

	order := TouristData{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// UpdateOrder - Updates an order
func (c *Client) UpdateOrder(orderID string, orderItems []OrderItem, authToken *string) (*Order, error) {
	rb, err := json.Marshal(orderItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/orders/%s", c.HostURL, orderID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	order := Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// DeleteOrder - Deletes an order
func (c *Client) DeleteOrder(orderID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/orders/%s", c.HostURL, orderID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "Deleted order" {
		return errors.New(string(body))
	}

	return nil
}
