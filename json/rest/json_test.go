package rest

import (
	"fmt"
)

const testString = `{
	"info" : {
		"id" : "0",
		"start" : 123123,
		"end" : 456456,
		"completed" : false,
		"dependencies" : ["3", "5"]
	},
	"tasks" : [
		{
			"id" : "1",
			"start" : 2319823,
			"end" : 4597979,
			"completed" : false,
			"dependencies" : ["3", "5"]
		},

		{
			"id" : "2", 
			"start" : 9357402,
			"end" : 29340572,
			"completed" : false,
			"dependencies" : ["3"]
		},

		{
			"id" : "3",
			"start" : 2542934,
			"end" : 2394575,
			"completed" : true,
			"dependencies" : []
		},

		{
			"id" : "4",
			"start" : 2934057,
			"end" : 1574234,
			"completed" : true,
			"dependencies" : ["1", "5"]
		},

		{
			"id" : "5",
			"start" : 21642323,
			"end" : 29345289,
			"completed" : false,
			"dependencies" : []
		}
	]
}`

func ExampleJSON() {
	data, err := UnmarshalResponse([]byte(testString))
	fmt.Println(data, err)
	str, err := MarshalResponse(data)
	fmt.Println(string(str), err)
	// Output: {{0 123123 456456 false [3 5]} [{1 2319823 4597979 false [3 5]} {2 9357402 29340572 false [3]} {3 2542934 2394575 true []} {4 2934057 1574234 true [1 5]} {5 21642323 29345289 false []}]} <nil>
	// {"info":{"id":"0","start":123123,"end":456456,"completed":false,"dependencies":["3","5"]},"tasks":[{"id":"1","start":2319823,"end":4597979,"completed":false,"dependencies":["3","5"]},{"id":"2","start":9357402,"end":29340572,"completed":false,"dependencies":["3"]},{"id":"3","start":2542934,"end":2394575,"completed":true,"dependencies":[]},{"id":"4","start":2934057,"end":1574234,"completed":true,"dependencies":["1","5"]},{"id":"5","start":21642323,"end":29345289,"completed":false,"dependencies":[]}]} <nil>
}
