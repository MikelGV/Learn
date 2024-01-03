package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
    Page int
    Fruits []string
}

// Only exported fields will be encoded/decoded in JSON. Fields must start with
// capital letters to be exported.
type response2 struct {
    Page int `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
    // Go built in support for JSON encoding and decoding.

    // Fist we'll look at encoding basic data types to JSON strings.
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    // Here are some for slices and maps, wich encodes to JSON arrays and obj
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // JSON package can automatically encode your custom data types. Only
    // includes exported fields in the encoded output and will by default
    // use those names as the JSON keys.
    res1D := &response1{
        Page: 1,
        Fruits: []string{"apple", "peach", "pear"},
    }
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    // Can use tags on struct field declarations to customize the encoded JSON
    // key names.
    res2D := &response2{
        Page: 1,
        Fruits: []string{"apple", "peach", "pear"},
    }
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // Decoding JSON data into Go values.
    byt := []byte(`{"num":6.13, "str":["a","b"]}`)

    // Provide a variable where JSON package can put the decoded data.
    var dat map[string]interface{}

    // Here is the actual decoding, and a check for associated errors.
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    // In order to use the values in the decoded map, we'll need to convert them
    // to appropiate type.
    num := dat["num"].(float64)
    fmt.Println(num)

    // Accessing nested data requires a series of conversions.
    strs := dat["str"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    // Can also decode JSON into custom data types. This has the advantage of
    // adding additional type-safety to programs eliminating the need for type
    // assertions when accessing the decoded data.
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[1])

    // In the example above we always used bytes and strings as intermediates
    // between the data and JSON representation on standard out. We can also
    // stream JSON encodings directly to os.Writers like os.Stdout or even HTTP
    // response bodies
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
