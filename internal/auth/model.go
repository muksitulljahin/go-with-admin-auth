package auth

// User ডাটাবেজ মডেল
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // JSON রেসপন্সে পাসওয়ার্ড হাইড থাকবে
	Role     string `json:"role"` // "admin" অথবা "user"
}

// LoginRequest ক্লায়েন্ট থেকে আসা ইনপুট ডেটা
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
