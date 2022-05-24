package jsonx

import "github.com/tidwall/gjson"

// Parse parses the json and returns a result.
func Parse(json string) gjson.Result {
	return gjson.Parse(json)

}

// Get searches json for the specified path.
func Get(json, path string) gjson.Result {
	return gjson.Get(json, path)
}
