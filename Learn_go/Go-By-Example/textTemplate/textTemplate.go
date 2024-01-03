package main

import (
	"os"
	"text/template"
)

func main() {

    // Templates are a mix of text and "actions" enclosed in {{.}} that are used
    // to dinamically insert content.
    t1 := template.New("t1")
    t1, err := t1.Parse("Value is {{.}}\n")
    if err != nil {
        panic(err)
    }

    // Alternatively use temple.Must to panic in case Parse returns error usefull
    // in the global scope.
    t1 = template.Must(t1.Parse("Value is {{.}}\n"))

    // Execute the template generate its text with specific values for its actions
    t1.Execute(os.Stdout, "some text")
    t1.Execute(os.Stdout, 5)
    t1.Execute(os.Stdout, []string{
        "Go",
        "Rust",
        "C++",
        "C#",
    })

    // Helper function
    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }

    // If data is a struct we can use {{.FieldName}} action to access its fields
    t2 := Create("t2", "Name: {{.Name}}\n")

    t2.Execute(os.Stdout, struct {
        Name string
    }{"Jane Doe"})

    // With maps these is no restriction on the case of key names
    t2.Execute(os.Stdout, map[string]string{
        "Name": "Mickey Mouse",
    })

    // If/Else provide conditional execution for templates. A value is considered
    // False if it's the default value type such as 0, an empty string, nil pointer
    // etc.
    t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
    t3.Execute(os.Stdout, "not empty") 
    t3.Execute(os.Stdout, "")

    // Range blocks let us loop through slices, arrays, maps or channels
    t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout, []string{ "Go", "Rust", "C++", "C#"})
}
