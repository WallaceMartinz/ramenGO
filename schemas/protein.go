package schemas

type Protein struct {
	ID            int     `json:"id"`
	ImageInactive string  `json:"imageInactive"`
	ImageActive   string  `json:"imageActive"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
}

var (
	proteins = []Protein{
		{
			ID: 			1,
			ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
			Name: 		   "Chasu",
			Description:   "A sliced flavourful pork meat with a selection of season vegetables.",
			Price: 		    10,
		},
		{
			ID: 			2,
			ImageInactive: "https://tech.redventures.com.br/icons/yasai/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/yasai/active.svg",
			Name: 		   "Yasai Vegetarian",
			Description:   "A delicious vegetarian lamen with a selection of season vegetables.",
			Price: 		    10,
		},
		{
			ID: 		    3,
			ImageInactive: "https://tech.redventures.com.br/icons/chicken/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/chicken/active.svg",
			Name: 		   "Karaague",
			Description:   "Three units of fried chicken, moyashi, ajitama egg and other vegetables.",
			Price: 			12,
		},
	}
)