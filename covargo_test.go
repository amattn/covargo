package covargo

import (
	"log"
	"testing"
)

func TestReadmeUsageDefaultExample(t *testing.T) {
	const (
		DROPBOOKSOFT_SECRET_APP_TOKEN = "DROPBOOKSOFT_SECRET_APP_TOKEN"
	)

	Add(DROPBOOKSOFT_SECRET_APP_TOKEN, "d", "dbs_token", "dbs_secret_app_token")

	Load(DROPBOOKSOFT_SECRET_APP_TOKEN)

	token := StringValue(DROPBOOKSOFT_SECRET_APP_TOKEN)

	log.Println("token:", token)
}

func TestReadmeUsageCustomExample(t *testing.T) {

	const (
		DROPBOOKSOFT_SECRET_APP_TOKEN = "DROPBOOKSOFT_SECRET_APP_TOKEN"
	)

	col := NewCollection()

	col.Add(DROPBOOKSOFT_SECRET_APP_TOKEN, "d", "dbs_token", "dbs_secret_app_token")

	col.Load(DROPBOOKSOFT_SECRET_APP_TOKEN)

	token := col.StringValue(DROPBOOKSOFT_SECRET_APP_TOKEN)

	// use token to connect to API, or whatever
	log.Println("token:", token)

}
