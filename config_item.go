package covargo

import "time"

// we use wonky casing to hint what the standard usage is.
type Item struct {
	ENV_VAR_NAME string // required
	Shortflag    string // short version of cli flag, such as -v. if "", we don't set a short flag
	Longflag     string // long version of cli flag, such as -version. if "", use ENV_VAR_NAME
	Json_key     string // lookup key for pulling value out of a json map.  if "", use ENV_VAR_NAME

	FlagUsage string // help message for cli flag(s)

	DefaultValue string
	RawValue     string // value we load from ENV, cli flag, file, etc.

	LastLoad time.Time

	// TODO
	//Type reflect.Type // TODO we default to strings for now, but eventually, will have type specific getters for ints, floats, etc.
	//IsSecret bool     // TODO we want to support exporting settings to a JSON file or something.  for secrets, we can't export them.
}

// load the raw value from cli flag, env var, file, etc.
func (ci Item) LoadValue() {

}

func (ci Item) String() string {
	return ci.RawValue
}
