structmaker
===========

A simple utility to generate struct from json for Go

## Usage
```sh
$ cat sample.json
{
	"dog_name": "bokdol",
	"dog_age": 5,
	"owner": {
		"name": "scryner"
	}

}

$ cat sample.json | structmaker Dog
type Dog struct {
	DogName string `json:"dog_name"`
	DogAge int `json:"dog_age"`
	Owner DogOwner `json:"owner"`
}

type DogOwner struct {
	Name string `json:"name"`
}
```
