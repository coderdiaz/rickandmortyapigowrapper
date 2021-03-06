package rickandmorty

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func GetCharacters(options map[string]interface{}) (*AllCharacters, error) {
	endpoint := endpointCharacter

	hasParams := false
	params := make(map[string]string)

	if options == nil {
		options = map[string]interface{}{
			"endpoint": endpoint,
		}
	}

	for k, v := range options {
		switch v.(type) {
		case int:
			if k == "page" {
				hasParams = true
				params[k] = strconv.FormatInt(int64(v.(int)), 10)
			}
			delete(options, k)
		case string:
			// Skip endpoint in options
			if k == "endpoint" {
				continue
			}
			// Valid parameters to be passed to the parameters map
			validParams := []string{"name", "status", "species", "type", "gender"}
			exists := containsString(validParams, k)
			if exists {
				hasParams = true
				params[k] = v.(string)
			}
			// Cleanup the options map
			delete(options, k)
		default:
			// Cleanup the options map
			delete(options, k)
			// Set the endpoint
			options["endpoint"] = endpoint
		}
	}

	if hasParams {
		options["endpoint"] = endpoint
		options["params"] = params
	}

	data, err := makePetition(options)
	if err != nil {
		return &AllCharacters{}, err
	}

	characters := new(AllCharacters)

	if err := mapstructure.Decode(data, &characters); err != nil {
		return &AllCharacters{}, err
	}

	return characters, nil
}

func GetCharacter(integer int) (*Character, error) {
	endpoint := endpointCharacter

	options := map[string]interface{}{
		"endpoint": endpoint,
		"params": map[string]int{
			"integer": integer,
		},
	}

	data, err := makePetition(options)
	if err != nil {
		return &Character{}, err
	}

	character := new(Character)

	if err := mapstructure.Decode(data, &character); err != nil {
		return &Character{}, err
	}

	return character, nil
}

func GetCharactersArray(integers []int) (*MultipleCharacters, error) {
	endpoint := endpointCharacter

	options := map[string]interface{}{
		"endpoint": endpoint,
		"integers": integers,
	}

	data, err := makePetition(options)
	if err != nil {
		return &MultipleCharacters{}, err
	}

	characters := new(MultipleCharacters)

	if err := mapstructure.Decode(data, &characters); err != nil {
		return &MultipleCharacters{}, err
	}

	return characters, nil
}
