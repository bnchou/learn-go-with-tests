package maps

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("definition already exist")
	ErrWordDoesNotExist = DictionaryErr("word doesn't exist")
)

type Dictionary map[string]string

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		return ErrWordExists
	case ErrNotFound:
		d[key] = value
		return nil
	default:
		return err
	}
}

func (d Dictionary) Search(s string) (string, error) {
	res, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}
	return res, nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}
