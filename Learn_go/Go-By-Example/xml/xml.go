package main

import (
	"encoding/xml"
	"fmt"
)

// Plant mapped to XML. Field tags contains directives for encoder and decoder.
// We use some special features of the XML package: XMLName field name dictates
// the name of the XML element. id,attr means that Id field is an XML attribute
type Plant struct {
    XMLName xml.Name `xml:"plant"`
    Id int `xml:"id,attr"`
    Name string `xml:"name"`
    Origin []string `xml:"origin"`
}

func (p Plant) String() string {
    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}


func main() {
    // Go offers build-in support for XML and XML-like formats

    coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}

    // Emit XML representing our plant; using MarshalIndent to produce a more
    // human-readable output.
    out, _ := xml.MarshalIndent(coffee, " ", "  ")
    fmt.Println(string(out))

    // To add a generic XML header to the output, append it explicitly
    fmt.Println(xml.Header + string(out))

    // Use Unmarshal to parse a stream of bytes with XML into a data structure
    // if not throw an error
    var p Plant
    if err := xml.Unmarshal(out, &p); err != nil {
        panic(err)
    }
    fmt.Println(p)

    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}

    // The parent>child>plant field tag tells the encoder to nest all plants
    // under parent>child>...

    type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants []*Plant `xml:"parent>child>plant"`
    }

    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}

    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))        
}
