package tourists3

// Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"data,omitempty"`
}

// OrderItem -
type OrderItem struct {
	Tourist   Tourists `json:"tourist"`
	//Tourist_Name      string         `json:"tourist_name"`
	//Tourist_Email     string         `json:"tourist_email"`
	//Tourist_Location  string         `json:"tourist_location"`
}

// Tourists -
type Tourists struct {
//	ID                string         `json:"id"`
	ID                int            `json:"id"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	Gender            string         `json:"gender"`
	Status            string         `json:"status"`
}

type TouristInput struct {
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	Gender            string         `json:"gender"`
	Status            string         `json:"status"`
}
