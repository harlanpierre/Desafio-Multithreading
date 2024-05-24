package api

import (
	"Desafio-Multithreading/structs"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func GetCep(ctx context.Context, url string, resultChan chan<- structs.Response) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Falha ao criar a requisição: %s", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Falha na requisição: %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Erro na resposta HTTP: %s", resp.Status)
		return
	}

	var address structs.Address
	if strings.Contains(url, "brasilapi") {
		var brasilApi structs.BrasilApi
		if err := json.NewDecoder(resp.Body).Decode(&brasilApi); err != nil {
			log.Printf("Falha ao converter JSON para BrasilAPI: %s", err)
			return
		}
		address = structs.Address{
			Cep:        brasilApi.Cep,
			Uf:         brasilApi.State,
			Localidade: brasilApi.City,
			Bairro:     brasilApi.Neighborhood,
			Logradouro: brasilApi.Street,
		}

		resultChan <- structs.Response{Address: address, Source: "Brasil API"}

	} else {
		var viaCep structs.ViaCep
		if err := json.NewDecoder(resp.Body).Decode(&viaCep); err != nil {
			log.Printf("Falha ao converter JSON para ViaCEP: %s", err)
			return
		}
		address = structs.Address{
			Cep:        viaCep.Cep,
			Uf:         viaCep.Uf,
			Localidade: viaCep.Localidade,
			Bairro:     viaCep.Bairro,
			Logradouro: viaCep.Logradouro,
		}

		resultChan <- structs.Response{Address: address, Source: "ViaCEP"}

	}
}
