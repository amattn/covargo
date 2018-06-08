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
3. cli flag or defalut filepath that points to single file
4. cli flag or defalut filepath that points to json file // TODO


Basic usage:

- Create a collection or use the default collection.
- Add congig Items, each identified by a string key.
- For each item, define how it can be loaded.
- Load the items.
- Each item can then be queried for its value.

Here's an example of using the default collection:

	const (
		DROPBOOKSOFT_SECRET_APP_TOKEN = "DROPBOOKSOFT_SECRET_APP_TOKEN"
	)

	item := covargo.Add(DROPBOOKSOFT_SECRET_APP_TOKEN)
	item.SetEnvVar(DROPBOOKSOFT_SECRET_APP_TOKEN)
	item.SetCliValueFlags("d", "dbs_token", "secret token used to access dbs API")

	// here we are using the same string for the item key as the env var.
	item.SetEnvVar(DROPBOOKSOFT_SECRET_APP_TOKEN)

	covargo.Load(DROPBOOKSOFT_SECRET_APP_TOKEN)

	token := covargo.StringValue(DROPBOOKSOFT_SECRET_APP_TOKEN)

	log.Println("token:", token)


Here's an example of using your own collection:

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
    

## Current Limitations

- Only supporting string values at the moment
- Flags used the built-in flagset, flag.CommandLine. 

## License

MIT.  Please see the LICENSE file for the usual boilerplate.