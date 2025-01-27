package requests

type RequestUpdateUser struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Smith"`
	Username  string `json:"username" example:"johnsmith"`
	Password  string `json:"password" example:"P@ssw0rd123"`
	Gender    string `json:"gender" example:"male"`
	Age       int    `json:"age" example:"28"`
	Birthday  string `json:"birthday" example:"1995-06-15T00:00:00Z"`
	Address   string `json:"address" example:"123 Main Street, Apt 4B"`
	City      string `json:"city" example:"San Francisco"`
	State     string `json:"state" example:"CA"`
	Country   string `json:"country" example:"USA"`
}

type RequestRegister struct {
	Email    string `json:"email" example:"john.smith@example.com"`
	Username string `json:"username" example:"johnsmith"`
	Password string `json:"password" example:"P@ssw0rd123"`
}

type RequestLogin struct {
	Email    string `json:"email" example:"john.smith@example.com"`
	Password string `json:"password" example:"P@ssw0rd123"`
}

type RequestLoginByUsername struct {
	Username string `json:"username" example:"johnsmith"`
	Password string `json:"password" example:"P@ssw0rd123"`
}
