package main

import (
	"Desafio-Multithreading/api"
	"Desafio-Multithreading/structs"
	"context"
	"fmt"
	"time"
)

const (
	timeoutDuration = 1 * time.Second
	cep             = "53180000"
	urlBrasilapi    = "https://brasilapi.com.br/api/cep/v1/" + cep
	urlViaCep       = "http://viacep.com.br/ws/" + cep + "/json/"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	resultChan := make(chan structs.Response, 2)

	go api.GetCep(ctx, urlViaCep, resultChan)
	go api.GetCep(ctx, urlBrasilapi, resultChan)

	select {
	case result := <-resultChan:
		fmt.Printf("API: %s\nResposta: %+v\n", result.Source, result.Address)
	case <-ctx.Done():
		fmt.Printf("Erro de timeout: %s\n", ctx.Err())
	}
}
