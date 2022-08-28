package maps

var testDefinition = "a procedure intended to establish the quality, performance, or reliability of something, especially before it is taken into widespread use."

type DictionaryErr string

const (
	errNoKey           = DictionaryErr("could not find key in dictionary")
	errExistingElement = DictionaryErr("cannot add existing element")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errNoKey
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNoKey:
		d[word] = definition
	case nil:
		return errExistingElement
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNoKey:
		return errNoKey
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
