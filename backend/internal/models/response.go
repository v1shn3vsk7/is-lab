package models

type EmptyResponse struct{}

type ErrResponse struct {
	Err string `json:"error"`
}
