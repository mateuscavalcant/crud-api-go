package request

type UserRequest struct {
	Name     string `json:"name" bidding:"required, min=3, max=50"`
	Age      int8   `json:"age" bidding:"required, min=10, max=120"`
	Email    string `json:"email" bidding:"required, email"`
	Password string `json:"password" bidding:"required, min=8, max=16"`
}
