package auth

import "errors"

// AuthService ইন্টারফেস
type Service interface {
	Login(req LoginRequest) (string, error)
}

type authService struct {
	// ভবিষ্যতে এখানে repository বা db যোগ হবে
}

// NewAuthService কনস্ট্রাকটর
func NewAuthService() Service {
	return &authService{}
}

func (s *authService) Login(req LoginRequest) (string, error) {
	// সিম্পল ডেমো বিজনেস লজিক
	if req.Email == "a@a.com" && req.Password == "123456" {
		// এখানে আসল সিস্টেমে JWT টোকেন রিটার্ন হবে
		return "fake-jwt-token-for-admin", nil
	}
	return "", errors.New("invalid email or password")
}
