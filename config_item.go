package covargo

import (
	"flag"
	"io/ioutil"
	"os"
	"time"

	"github.com/amattn/deeperror"
)

type LoadMethod int

const (
	LongFlag LoadMethod = iota
	ShortFlag
	EnvVar
	FileContentsLongFlag
	FileContentsShortFlag
	FileContentsDefaultLocation

	JsonFileLocationLongFlag
	JsonFileLocationShortFlag
	JsonFileDefaultLocation

	Error
	LoadMethodCount
)

// we use wonky casing to hint what the standard usage is.
type Item struct {
	Key string // required

	// method cli flag
	Longflag                  string // long version of cli flag, such as -version. if "", we don't set a long flag
	Shortflag                 string // short version of cli flag, such as -v. if "", we don't set a short flag
	FlagUsage                 string // help message for cli flag(s)
	cli_flag_short_string_ptr *string
	cli_flag_long_string_ptr  *string

	// method: environment variable value
	ENV_VAR_NAME string // environment variable name. if "" this variable cannot be loaded from environment variable

	// method: read a file, use contents of file as value
	FileContentsPathLongFlag       string // a cli flag that tells where to load a file and use the entire contents of the file as the config variable.
	FileContentsPathShortFlag      string // a cli flag that tells where to load a file and use the entire contents of the file as the config variable.
	FileContentsDefaultPath        string // default value to look for file to read and use entire contents as value of config var
	FileContentsUsage              string
	file_contents_short_string_ptr *string
	file_contents_long_string_ptr  *string

	// method Json file
	JSONFilePathLongFlag  string
	JSONFilePathShortFlag string
	JSONFileDefaultPath   string
	Json_key              string // lookup key for pulling value out of a json map.  if "", this variable cannot be loaded from a json file

	DefaultValue string // none of the methods work? use this default value

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
func NewItem(key string) *Item {
	item := new(Item)
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
		ci.cli_flag_short_string_ptr = flag.String(ci.Shortflag, ci.DefaultValue, ci.FlagUsage)
	}
	if longflag != "" {
		ci.cli_flag_long_string_ptr = flag.String(ci.Longflag, ci.DefaultValue, ci.FlagUsage)
	}
}

func (ci *Item) SetFileContentsFlags(shortflag, longflag, default_path, usage string) {
	ci.FileContentsPathLongFlag = longflag
	ci.FileContentsPathShortFlag = shortflag
	ci.FileContentsDefaultPath = default_path

	if shortflag != "" {
		ci.file_contents_short_string_ptr = flag.String(shortflag, default_path, usage)
	}
	if longflag != "" {
		ci.file_contents_long_string_ptr = flag.String(longflag, default_path, usage)
	}
}

// load the raw value from cli flag, env var, file, etc.
// returns true if
func (ci *Item) LoadValue() error {
	if len(ci.Shortflag) > 0 || len(ci.Longflag) > 0 {
		if flag.Parsed() == false {
			return deeperror.New(1254443215, "flag.Parsed == false.  cannot load value with unparsed flags", nil)
		}

		if len(ci.Longflag) > 0 && ci.cli_flag_long_string_ptr != nil {
			ci.set_raw_value(*ci.cli_flag_long_string_ptr, LongFlag)
			return nil
		}

		if len(ci.Shortflag) > 0 && ci.cli_flag_short_string_ptr != nil {
			ci.set_raw_value(*ci.cli_flag_short_string_ptr, ShortFlag)
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

	if ci.FileContentsPathLongFlag != "" ||
		ci.FileContentsPathShortFlag != "" ||
		ci.FileContentsDefaultPath != "" {
		if flag.Parsed() == false {
			return deeperror.New(1143790062, "flag.Parsed == false.  cannot load value with unparsed flags", nil)
		}

		if ci.FileContentsPathLongFlag != "" && ci.file_contents_long_string_ptr != nil {
			filepath := *ci.file_contents_long_string_ptr
			raw_bytes, err := ioutil.ReadFile(filepath)
			if err != nil {
				// ignore? log?
			} else {
				ci.set_raw_value(string(raw_bytes), FileContentsLongFlag)
				return nil
			}
		}

		if ci.FileContentsPathShortFlag != "" && ci.file_contents_short_string_ptr != nil {
			filepath := *ci.file_contents_short_string_ptr
			raw_bytes, err := ioutil.ReadFile(filepath)
			if err != nil {
				// ignore? log?
			} else {
				ci.set_raw_value(string(raw_bytes), FileContentsShortFlag)
				return nil
			}
		}

		if ci.FileContentsDefaultPath != "" {
			filepath := ci.FileContentsDefaultPath
			raw_bytes, err := ioutil.ReadFile(filepath)
			if err != nil {
				// ignore? log?
			} else {
				ci.set_raw_value(string(raw_bytes), FileContentsDefaultLocation)
				return nil
			}
		}
	}

	// TODO: other methods

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
