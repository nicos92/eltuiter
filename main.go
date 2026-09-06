package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/nicos92/eltuiter/awsgo"
	"github.com/nicos92/eltuiter/bd"
	"github.com/nicos92/eltuiter/handlers"
	"github.com/nicos92/eltuiter/models"
	"github.com/nicos92/eltuiter/secretmanager"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InicioAWS()

	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "error en las variables de entorno. deben incluir 'secretname', 'bucketname', 'urlprefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("secretname"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "error en la lectura de Secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	fmt.Println("Chequeo agregar todo al contexto")

	pathParam, ok := request.PathParameters["eltuiter-resource"]
	if !ok || pathParam == "" {
		// Si PathParameters es nil o no trae la clave, usa la ruta completa directa
		pathParam = request.Path
	}
	path := strings.ReplaceAll(request.PathParameters["eltuiter-resource"], os.Getenv("urlprefix"), "")

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketname"), os.Getenv("bucketname"))

	// chequeo conexión con la BD
	//

	fmt.Println("Chequeo la conexion con la base de datos. - El path: " + path)
	err = bd.ContarBD(awsgo.Ctx)

	if err != nil {
		fmt.Print("error al conectar la BD" + err.Error())
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "error conectando a la BD" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	fmt.Println("Configuro los handlers")
	resApi := handlers.Manejadores(awsgo.Ctx, request)
	if resApi.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: resApi.Status,
			Body:       resApi.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {

		return resApi.CustomResp, nil
	}
}

func ValidoParametros() bool {
	_, SecretName := os.LookupEnv("secretname")
	if !SecretName {
		return SecretName
	}
	_, BucketName := os.LookupEnv("bucketname")
	if !BucketName {
		return BucketName
	}
	_, UrlPrefix := os.LookupEnv("urlprefix")
	if !UrlPrefix {
		return UrlPrefix
	}

	return true
}
