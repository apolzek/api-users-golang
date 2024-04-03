package viacep

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wiliamvj/api-users-golang/config/env"
)

func TestGetCep(t *testing.T) {
	env.InitConfig()
	env.Env.ViaCepURL = "https://viacep.com.br/ws"

	expected := &ViaCepResponse{
		CEP:         "38702-122",
		Logradouro:  "Rua Maria Borges Silva",
		Complemento: "",
		Bairro:      "Abner Afonso",
		Localidade:  "Patos de Minas",
		UF:          "MG",
		IBGE:        "3148004",
		GIA:         "",
		DDD:         "34",
		SIAFI:       "4959",
	}

	result, err := GetCep("38702122")
  // t.Logf("Result: %+v", result)
	assert.NoError(t, err, "GetCep should not receive an error on the call")

	assert.Equal(t, expected, result, "The result of GetCep must be equal to expected")
}
