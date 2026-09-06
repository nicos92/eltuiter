package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nicos92/eltuiter/jwt"
	"github.com/nicos92/eltuiter/models"
	"github.com/nicos92/eltuiter/routers"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RestApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var res models.RestApi
	res.Status = 400

	idOk, statusCode, msg, claim := validoAuthorization(ctx, request)

	if !idOk {
		res.Status = statusCode
		res.Message = msg
		return res
	}
	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELELE":
		switch ctx.Value(models.Key("path")).(string) {

		}

	}

	res.Message = "Method Invalid"
	return res
}

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	switch path {
	case "registro", "login", "obtenerAvatar", "obtenerBanner":
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "token requerido", models.Claim{}
	}

	claim, todoOK, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtsign")).(string))

	if !todoOK {
		if err != nil {
			fmt.Println("error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		}
		fmt.Println("error en el token " + msg)
		return false, 401, msg, models.Claim{}

	}

	fmt.Println("Token OK")
	return true, 200, msg, *claim
}
