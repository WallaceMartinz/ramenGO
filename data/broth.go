package data

import (
	"errors"
	"strconv"
)

// Broth represents a type of broth available in the application.
type Broth struct {
	ID            int     `json:"id"`
	ImageInactive string  `json:"imageInactive"`
	ImageActive   string  `json:"imageActive"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
}

// broths is a slice containing predefined broth data.
var broths = []Broth{
	{
		ID:            1,
		ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/salt/active.svg",
		Name:          "Salt",
		Description:   "Simple like the seawater, nothing more",
		Price:         10,
	},
	{
		ID:            2,
		ImageInactive: "https://tech.redventures.com.br/icons/shoyu/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/shoyu/active.svg",
		Name:          "Shoyu",
		Description:   "The good old and traditional soy sauce",
		Price:         10,
	},
	{
		ID:            3,
		ImageInactive: "https://tech.redventures.com.br/icons/miso/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/miso/active.svg",
		Name:          "Miso",
		Description:   "Paste made of fermented soybeans",
		Price:         12,
	},
}

// GetBroths retrieves a list of available broths.
func GetBroths() ([]Broth, error) {
	if len(broths) == 0 {
		return nil, errors.New("no broths available")
	}
	return broths, nil
}

// GetBrothNameById returns the name of the broth based on the provided ID.
func GetBrothNameById(brothId string) (string, error) {
	id, err := strconv.Atoi(brothId)
	if err != nil {
		return "", err
	}

	broths, err := GetBroths()
	if err != nil {
		return "", err
	}

	for _, broth := range broths {
		if broth.ID == id {
			return broth.Name, nil
		}
	}
	return "", errors.New("invalid broth ID")
}
