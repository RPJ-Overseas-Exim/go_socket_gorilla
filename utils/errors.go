package utils

type HTTPException struct {
    Message string
}

func (he *HTTPException) Error() string{
    return  he.Message
}

