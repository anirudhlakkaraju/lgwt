package maps

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("could not find word you were looking for")
	ErrWordExists = DictionaryErr("word already exists in dictionary")
)

type DictionaryErr string

// Implements std error interface
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) {
	d[key] = value
}
