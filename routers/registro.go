package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nicos92/eltuiter/bd"
	"github.com/nicos92/eltuiter/models"
)

func Registro(ctx context.Context) models.RestApi {
	var usuario models.Usuario
	var restApi models.RestApi

	restApi.Status = 400

	fmt.Println("Entré al Registro")

	body := ctx.Value(models.Key("body")).(string)

	err := json.Unmarshal([]byte(body), &usuario)
	if err != nil {
		restApi.Message = err.Error()
		fmt.Println(restApi.Message)
		return restApi
	}
	if len(usuario.Email) == 0 {
		restApi.Message = "debe especificar el email"
		fmt.Println(restApi.Message)

		return restApi
	}
	if len(usuario.Password) == 0 {
		restApi.Message = "debe especificar el password de al menos 6 caracteres"
		fmt.Println(restApi.Message)

		return restApi
	}

	_, existe, _ := bd.ChequeoYaExisteUsuario(usuario.Email)

	if existe {
		restApi.Message = "ya existe un usuario registrado con ese email"
		fmt.Println(restApi.Message)
		return restApi
	}

	_, status, err := bd.InsertoRegistro(usuario)

	if err != nil {
		restApi.Message = "ocurrió un error al intertar realizar el registro del usuario " + err.Error()
		fmt.Println(restApi.Message)
		return restApi
	}

	if !status {
		restApi.Message = "no se ha logrado insertar el registro del usuario"
		fmt.Println(restApi.Message)
		return restApi
	}

	restApi.Status = 200
	restApi.Message = "Registro OK"
	fmt.Println(restApi.Message)
	return restApi
}
