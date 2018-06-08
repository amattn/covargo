package covargo

type Collection map[string]*Item

func NewCollection() Collection {
	return Collection{}
}

// add a config item to our "library" of items
// will panic if ENV_VAR_NAME is empty
// CliFlag and JSONKey can be empty, and will default to ENV_VAR_NAME
// typically, CliFlag is shorter and easy to type and JSONKey is empty or a lowercase, snake case variant of ENV_VAR_NAME
func (col Collection) Add(key string) *Item {
	ci := NewItem(key)
	col[key] = ci
	return ci
}

func (col Collection) Get(key string) *Item {
	return col[key]
}

func (col Collection) Load(key string) error {
	return col.Get(key).LoadValue()
}

func (col Collection) LoadAll() error {
	for _, item := range col {
		err := item.LoadValue()
		if err != nil {
			// TODO use multierr or similar to coallate all errors and not block
			return err
		}
	}
	return nil
}

func (col Collection) Contains(key string) bool {
	_, exists_in_map := col[key]
	return exists_in_map
}

func (col Collection) Remove(key string) {
	delete(col, key)
}

// convenience
func (col Collection) StringValue(key string) string {
	return col.Get(key).StringValue()
}
