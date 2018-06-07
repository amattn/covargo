# covargo

Config Variables helper library in Go.

Go (golang) library for loading configuration variables from cli flags, env vars or files

## Installation

Same as most go libraries:

    go get -u https://github.com/amattn/covargo

or if you are into vgo, just add the `import ` statement and do a `vgo build`

## Usage

`covargo` makes it easy to load configuration variables in a variety of ways.  

in order of highest priority first:

1. Command line flag
2. Environment variable
3. Entire Contents of Single file
4. cli flag that points to single file
5. Contents of a json map in a file // TODO
6. cli flag that points to json file // TODO


Basic usage:

- Create a collection or use the default collection
- 

Here's an example of using the default collection:

	const (
		DROPBOOKSOFT_SECRET_APP_TOKEN = "DROPBOOKSOFT_SECRET_APP_TOKEN"
	)

	covargo.Add(DROPBOOKSOFT_SECRET_APP_TOKEN, "d", "dbs_token", "dbs_secret_app_token")

	covargo.Load(DROPBOOKSOFT_SECRET_APP_TOKEN)

	token := covargo.StringValue(DROPBOOKSOFT_SECRET_APP_TOKEN)

	log.Println("token:", token)


Here's an example of using your own collection:

	const (
		DROPBOOKSOFT_SECRET_APP_TOKEN = "DROPBOOKSOFT_SECRET_APP_TOKEN"
	)

	col := NewCollection()

	col.Add(DROPBOOKSOFT_SECRET_APP_TOKEN, "d", "dbs_token", "dbs_secret_app_token")

	col.Load(DROPBOOKSOFT_SECRET_APP_TOKEN)

	token := col.StringValue(DROPBOOKSOFT_SECRET_APP_TOKEN)

	// use token to connect to API, or whatever
	log.Println("token:", token)
    

## License

MIT.  Please see the LICENSE file for the usual boilerplate.