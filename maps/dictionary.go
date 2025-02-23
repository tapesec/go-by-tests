package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryError("could not find the word you were looking for")
	ErrWordExists       = DictionaryError("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryError("word doesn't exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(w string) (string, error) {
	definition, ok := d[w]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(w, definition string) error {
	_, err := d.Search(w)
	switch err {
	case nil:
		return ErrWordExists
	case ErrNotFound:
		d[w] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(w, definition string) error {
	_, err := d.Search(w)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[w] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(w string) error {
	_, err := d.Search(w)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, w)
	default:
		return err
	}
	return nil

}
