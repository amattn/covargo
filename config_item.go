package covargo

import (
	"flag"
	"os"
	"time"

	"github.com/amattn/deeperror"
)

type LoadMethod int

const (
	ShortFlag LoadMethod = iota
	LongFlag
	EnvVar
	WholeFileValue
	JsonFile
	Error
	LoadMethodCount
)

// we use wonky casing to hint what the standard usage is.
type Item struct {
	Key          string // required
	Shortflag    string // short version of cli flag, such as -v. if "", we don't set a short flag
	Longflag     string // long version of cli flag, such as -version. if "", we don't set a long flag
	ENV_VAR_NAME string // environment variable name. if "" this variable cannot be loaded from environment variable

	SingleValueFilePathFlag      string // a cli flag that tells where to load a file and use the entire contents of the file as the config variable.
	SingleValueFileShortPathFlag string // a cli flag that tells where to load a file and use the entire contents of the file as the config variable.
	SingleValueFileDefaultPath   string // default value to look for file to read and use entire contents as value of config var

	JSONFilePathFlag      string
	JSONFileShortPathFlag string
	JSONFileDefaultPath   string
	Json_key              string // lookup key for pulling value out of a json map.  if "", this variable cannot be loaded from a json file

	FlagUsage              string // help message for cli flag(s)
	DefaultValue           string
	short_string_ptr       *string
	long_string_ptr        *string
	single_file_string_ptr *string

	RawValue string // value we load from ENV, cli flag, file, etc.

	LoadMethod LoadMethod // how was the value loaded?
	LastLoad   time.Time  // when was it loaded?

	// TODO
	//Type reflect.Type // TODO we default to strings for now, but eventually, will have type specific getters for ints, floats, etc.
	//IsSecret bool     // TODO we want to support exporting settings to a JSON file or something.  for secrets, we can't export them.
}

// add a config item to our "library" of items
// will panic if `key` is empty
// if ENV_VAR_NAME is "", then LoadValue() won't check environment variables.
// if Shortflag and/or longflag is "", LoadValue() won't check cli flags.
// if JSONKey is "", LoadValue() won't check json file.
// typically, environment variables are all caps and words are separated by underscores (_)
// typically, CliFlag is shorter and easy to type and JSONKey is empty or a lowercase, snake case variant of ENV_VAR_NAME
func MakeItem(key string) Item {
	item := Item{}
	item.Key = key
	return item
}

func (ci *Item) SetEnvVar(name string) {
	ci.ENV_VAR_NAME = name
}

func (ci *Item) SetDefaultValue(default_value string) {
	ci.DefaultValue = default_value
}

// for now, we only support strings...
func (ci *Item) SetCliValueFlags(shortflag, longflag, usage string) {
	ci.Shortflag = shortflag
	ci.Longflag = longflag
	ci.FlagUsage = usage

	if shortflag != "" {
		ci.short_string_ptr = flag.String(ci.Shortflag, ci.DefaultValue, ci.FlagUsage)
	}
	if longflag != "" {
		ci.long_string_ptr = flag.String(ci.Longflag, ci.DefaultValue, ci.FlagUsage)
	}
}

func (ci *Item) SetSingleFileContent(shortflag, longflag, default_path, usage string) {
	ci.SingleValueFilePathFlag = shortflag
	ci.SingleValueFileShortPathFlag = longflag
	ci.SingleValueFileDefaultPath = default_path

	if shortflag != "" {
		ci.single_file_string_ptr = flag.String(shortflag, default_path, usage)
	}
	if longflag != "" {
		ci.single_file_string_ptr = flag.String(longflag, default_path, usage)
	}
}

// load the raw value from cli flag, env var, file, etc.
// returns true if
func (ci *Item) LoadValue() error {
	if len(ci.Shortflag) > 0 || len(ci.Longflag) > 0 {
		if flag.Parsed() == false {
			return deeperror.New(1254443215, "flag.Parsed == false.  cannot load value with unparsed flags", nil)
		}

		if len(ci.Shortflag) > 0 {
			ci.set_raw_value(*ci.short_string_ptr, ShortFlag)
			return nil
		}

		if len(ci.Longflag) > 0 {
			ci.set_raw_value(*ci.long_string_ptr, LongFlag)
			return nil
		}
	}

	if len(ci.ENV_VAR_NAME) > 0 {

		candidate_value, value_exists := os.LookupEnv(ci.ENV_VAR_NAME)

		if value_exists {
			// we are good here.  set our raw value and move on.
			ci.set_raw_value(candidate_value, EnvVar)
			return nil
		}
		// if value_exists == false, then fall thru and move along
	}

	// if we get here then that means none of the above worked...

	if ci.ENV_VAR_NAME == "" &&
		ci.Shortflag == "" &&
		ci.Longflag == "" &&
		ci.Json_key == "" {
		// item is misconfigured.  you need to set at least one of them
		return deeperror.New(2021170454, "Misconfigured covargo item.  No method to load configuration variable specified", nil)
	}

	// unable to load config variable.  for now, users must rely on the default variable
	return nil
}

func (ci *Item) set_raw_value(raw string, how LoadMethod) {
	ci.RawValue = raw
	ci.LoadMethod = how
	ci.LastLoad = time.Now()
}

func (ci Item) String() string {
	return ci.RawValue
}
