package utils

import (
	"log"
	"github.com/aidarkhanov/nanoid"
)

func GenNanoid() *string{
    alphabet := nanoid.DefaultAlphabet
    id, err := nanoid.Generate(alphabet, 12)
    if err!=nil{
        log.Fatalln("Nanoid Could not be generate: ", err)
    }
    return &id
}

