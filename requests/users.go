package requests

type RequestUser struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Smith"`
	Email     string `json:"email" example:"john.smith@example.com"`
	Username  string `json:"username" example:"johnsmith"`
	Password  string `json:"password" example:"P@ssw0rd123"`
	Gender    string `json:"gender" example:"male"`
	Age       uint8  `json:"age" example:"28"`
	Birthday  string `json:"birthday" example:"1995-06-15T00:00:00Z"`
	Address   string `json:"address" example:"123 Main Street, Apt 4B"`
	City      string `json:"city" example:"San Francisco"`
	State     string `json:"state" example:"CA"`
	Country   string `json:"country" example:"USA"`
}
