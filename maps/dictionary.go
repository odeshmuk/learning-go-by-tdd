package maps

import "errors"

//A type created for the map of string,string
type Dictionary map[string]string

//defining an error type for errors not found
var ErrNotFound = errors.New("could not find the word you were searching")

//defining an error type for an already existing key in dictionary
var ErrAlreadyExists = errors.New("already existing")

//defining that the key doesn't exist
var ErrKeyDoesntExist = errors.New("key doesn't exist")

//returns the result,error as part of this func
func (dictionary Dictionary) Search(key string) (string, error) {
	definition, ok := dictionary[key]
	//if not ok, return the group below
	if !ok {
		return "", ErrNotFound
	}
	//return this if a-ok
	return definition, nil
}

func (dictionary Dictionary) Add(key, value string) error {
	_, err := dictionary.Search(key)
	switch err {
	case ErrNotFound:
		dictionary[key] = value
	case nil:
		return ErrAlreadyExists
	}
	return nil

}

func (dictionary Dictionary) Update(key, updatedValue string) error {
	_, err := dictionary.Search(key)
	switch err {
	case ErrNotFound:
		return ErrKeyDoesntExist
	case nil:
		dictionary[key] = updatedValue
	}
	return nil
}

func (dictionary Dictionary) Delete(key string) error {
	_, err := dictionary.Search(key)
	switch err {
	case ErrNotFound:
		return ErrKeyDoesntExist
	case nil:
		delete(dictionary, key)
	}
	return nil
}
