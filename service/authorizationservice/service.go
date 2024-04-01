package authorizationservice

type Repository interface {
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
func (s Service) CheckAccess(userRole string, role string) bool {
	if userRole != role {
		return false
	}

	return true
}
