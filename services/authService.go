package services

type authService struct{
    
}

func NewAuthService() *authService{
    return &authService{
    }
}

func (as *authService) VerifyUser(username, password string) bool {
    adminUsername := "Rashid" 
    adminPassword := "Gp@12345"

    return adminUsername == username && adminPassword == password 
}
