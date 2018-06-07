package covargo

import (
	"log"
	"testing"
)

func TestReadmeUsageDefaultExample(t *testing.T) {
	const (
		DROPBOOKSOFT_SECRET_APP_TOKEN = "DROPBOOKSOFT_SECRET_APP_TOKEN"
	)

	item := Add(DROPBOOKSOFT_SECRET_APP_TOKEN)

	// here we are using the same string for the item key as the env var.
	item.SetEnvVar(DROPBOOKSOFT_SECRET_APP_TOKEN)

	Load(DROPBOOKSOFT_SECRET_APP_TOKEN)

	token := StringValue(DROPBOOKSOFT_SECRET_APP_TOKEN)

	log.Println("token:", token)
}

func TestReadmeUsageCustomExample(t *testing.T) {

	const (
		DROPBOOKSOFT_SECRET_APP_TOKEN = "DROPBOOKSOFT_SECRET_APP_TOKEN"
	)

	col := NewCollection()

	item := col.Add(DROPBOOKSOFT_SECRET_APP_TOKEN)

	item.SetEnvVar(DROPBOOKSOFT_SECRET_APP_TOKEN)
	item.SetCliValueFlags("d", "dbs_token", "secret token used to access dbs API")

	col.Load(DROPBOOKSOFT_SECRET_APP_TOKEN)

	token := col.StringValue(DROPBOOKSOFT_SECRET_APP_TOKEN)

	// use token to connect to API, or whatever
	log.Println("token:", token)

}
