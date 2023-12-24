package util

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func GenerateRequestId(urlPath string) string {

	valor := strconv.FormatInt(time.Now().UnixMilli(), 10) + "_" +  uuid.New().String()

	h := sha1.New()
	_, err := io.WriteString(h, valor)
	if  err != nil {
		fmt.Println("!! Erro", err)
		valor = valor + "_err"
	} else {
		valor = base64.StdEncoding.EncodeToString( h.Sum(nil) )
	}

	fmt.Println(">> VALOR: ", valor)

	return valor
}