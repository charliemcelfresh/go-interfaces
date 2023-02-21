package main

type app struct{}

type UserRequest struct {
	ID string `json:"id"`
}

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (a app) GetUser(u UserRequest) UserResponse {
	return UserResponse{
		ID:   u.ID,
		Name: "Charlie",
	}
}
