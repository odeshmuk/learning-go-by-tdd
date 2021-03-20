package maps

import "errors"

//A type created for the map of string,string
type Dictionary map[string]string

//defining an error type for errors not found
var ErrNotFound = errors.New("could not find the word you were searching")

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
