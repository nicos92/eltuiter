package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nicos92/eltuiter/models"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RestApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var res models.RestApi
	res.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

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
