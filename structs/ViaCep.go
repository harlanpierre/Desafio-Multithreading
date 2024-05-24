package structs

type ViaCep struct {
	Cep        string `json:"cep"`
	Uf         string `json:"uf"`
	Localidade string `json:"localidade"`
	Bairro     string `json:"bairro"`
	Logradouro string `json:"logradouro"`
}
