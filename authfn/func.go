package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"

	fdk "github.com/fnproject/fdk-go"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"github.com/oracle/oci-go-sdk/v65/secrets"
)

type Input struct {
	Type string `json:"type"`
	Token string `json:"token"`
}

type Output struct {
	Active bool `json:"active"`
}

func main() {
  fdk.Handle(fdk.HandlerFunc(myHandler))
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) { 

	config := fdk.GetContext(ctx).Config()
	secretid := config["secretid"]
	secret := getSecret(secretid)

	var msg Output
	var input Input

	json.NewDecoder(in).Decode(&input)

	if(secret == input.Token){
		msg.Active = true
	}else{
		msg.Active = false
	}

	json.NewEncoder(out).Encode(&msg)

}

func getSecret(SecretID string)(string){
	
	provider, err := auth.ResourcePrincipalConfigurationProvider()
	if err != nil {
		log.Println("Erro ao gerar o Provider")
		panic(err)
	}

	client, err := secrets.NewSecretsClientWithConfigurationProvider(provider)
	if err != nil {
		log.Println("Erro ao gerar o Client")
		panic(err)
	}
	req := secrets.GetSecretBundleRequest{Stage: secrets.GetSecretBundleStageCurrent, SecretId: common.String(SecretID)}

	resp, err := client.GetSecretBundle(context.Background(), req)
	if err != nil {
		log.Println("Erro Get Secret")
		panic(err)
	}

	base64Details := resp.SecretBundleContent.(secrets.Base64SecretBundleContentDetails)
	decodedContent, err := base64.StdEncoding.DecodeString(*base64Details.Content)
	if err != nil {
		log.Println("Erro ao Converter de base64")
		panic(err)
	}

	return string(decodedContent)
}