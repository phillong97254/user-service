package profile

type GetProfileInput struct {
	ID string
}

type GetProfileOutput struct {
	ID         string
	Name       string
	Email      string
	Department string
}
