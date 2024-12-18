package services

type authService struct{
    
}

func NewAuthService() *authService{
    return &authService{
    }
}

func (as *authService) VerifyUser(username, password string) bool {
    adminEmail := "abc@gmail.com" 
    adminPassword := "Gp@12345"

    return adminEmail == username && adminPassword == password 
}
