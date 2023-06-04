package tourists

// Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"data,omitempty"`
}

// OrderItem -
type OrderItem struct {
	Tourist   TouristData `json:"tourist"`
	//Tourist_Name      string         `json:"tourist_name"`
	//Tourist_Email     string         `json:"tourist_email"`
	//Tourist_Location  string         `json:"tourist_location"`
}

// Tourist -
type Tourist struct {
	Page              int            `json:"page"`
	Per_Page          int            `json:"per_page"`
	TotalRecord       int            `json:"totalrecord"`
	Total_Pages       int            `json:"total_pages"`
	Data              []TouristData  `json:"data"`
}

type TouristData struct {
	ID                string         `json:"id"`
	Tourist_Name      string         `json:"tourist_name"`
	Tourist_Email     string         `json:"tourist_email"`
	Tourist_Location  string         `json:"tourist_location"`
	Createdat         string         `json:"Createdat"`
}

type TouristInput struct {
	Tourist_Name      string         `json:"tourist_name"`
	Tourist_Email     string         `json:"tourist_email"`
	Tourist_Location  string         `json:"tourist_location"`
}
