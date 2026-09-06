package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InicioAWS() {
	Ctx = context.TODO()
	// TODO cambio el hardcodeo por variable de entorno
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		panic("error al cargar la configuracion .aws/config " + err.Error())
	}
}
