package maps

const (
	ErrNotFound          = DictionaryErr("could not found the word you're looking for")
	ErrWordExists        = DictionaryErr("word already exists")
	ErrWordDoesNotExists = DictionaryErr("word does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary store key value for words
type Dictionary map[string]string

// Search return words of key
func (d Dictionary) Search(key string) (string, error) {
	word := d[key]
	if word == "" {
		return "", ErrNotFound
	}

	return word, nil
}

// Add adds definition of word to Dictionary without modify existing word
// map is a `refference type` like a pointer
// We can modify map without specifying the pointer type
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	// default return other type of error
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newdef string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = newdef
	case ErrNotFound:
		return ErrWordDoesNotExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
