package steps


// MapperStep is used to map a maps field to another one.
type MapperStep struct {
	Name      string
	MapConfig map[string]string
}


// Build creates the "new" map with the new keys.
func (mapper MapperStep) Build() map[string]any {
	return nil
} 


// AddData takes the newly made map and adds the data to it.
func (mapper MapperStep) AddData(data map[string]any, newData map[string]any) map[string]any {
	return nil
}


// Process takes the map and maps the fields to a new map.
// First it checks if the map is not nil.
// Then it looks over MapConfig and and builds a new map with the new fields.
// Finally it takes the given data and puts it in the new app.
func (mapper MapperStep) Process(data any) (any, error) {
	return nil, nil
}
