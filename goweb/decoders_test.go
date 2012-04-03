package goweb

import (
	"testing"
	"net/url"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"strconv"
	"reflect"
)

type personTestStruct struct {
	Name      string
	Age       int
	Atoms     int64
	Nicknames []string
}

var personName string = "Alice"
var personAge int = 25
var personAtoms int64 = 29357029322375092
var personNicknames []string = []string{"ally", "al"}

func makeFormData() string {
	form := make(url.Values)
	form.Add("Name", personName)
	form.Add("Age", strconv.Itoa(personAge))
	form.Add("Atoms", strconv.FormatInt(personAtoms, 10))
	for _, name := range personNicknames {
		form.Add("Nicknames", name)
	}
	return form.Encode()
}

func makeTestContextWithContentTypeAndBody(ct, body string) *Context {
	var request *http.Request = new(http.Request)
	request.URL, _ = url.Parse("http://www.example.com/test?context=123")
	// add content type
	request.Header = make(http.Header)
	request.Header.Add("Content-Type", ct)
	// add form data as ReadCloser
	data := ioutil.NopCloser(bytes.NewBufferString(body))
	request.Body = data
	request.ContentLength = int64(len(body))
	request.Method = "POST"
	// setup context
	var responseWriter http.ResponseWriter
	var pathParams ParameterValueMap = make(ParameterValueMap)
	return makeContext(request, responseWriter, pathParams)
}

func TestFormDecoding(t *testing.T) {
	cx := makeTestContextWithContentTypeAndBody(
		"application/x-www-form-urlencoded; charset=utf8", makeFormData())
	// fill struct 
	var person personTestStruct
	err := cx.Fill(&person)
	if err != nil {
		t.Errorf("form-decoder:", err)
	}
	// check it
	if person.Name != personName {
		t.Errorf("form-decoder: expected %v got %v", personName, person.Name)
	}
	if person.Age != personAge {
		t.Errorf("form-decoder: expected %v got %v", personAge, person.Age)
	}
	if person.Atoms != personAtoms {
		t.Errorf("form-decoders: expected %v got %v", personAtoms, person.Atoms)
	}
	if !reflect.DeepEqual(person.Nicknames, personNicknames) {
		t.Errorf("form-decoders: expected %v got %v", personNicknames, person.Nicknames)
	}
}

func TestFormDecodingPtrPtr(t *testing.T) {
	cx := makeTestContextWithContentTypeAndBody(
		"application/x-www-form-urlencoded; charset=utf8", makeFormData())
	// fill struct via ptr -> ptr
	person := &personTestStruct{}
	err := cx.Fill(&person)
	if err != nil {
		t.Errorf("form-decoder:", err)
	}
	// check it
	if person.Name != personName {
		t.Errorf("form-decoder: expected %s got %v", personName, person.Name)
	}
	if person.Age != personAge {
		t.Errorf("form-decoder: expected %v got %v", personAge, person.Age)
	}
	if person.Atoms != personAtoms {
		t.Errorf("form-decoders: expected %v got %v", personAtoms, person.Atoms)
	}
}

func makeJsonData() string {
	return fmt.Sprintf(`{"Name":"%s", "Age":%d, "Atoms":%d, "Nicknames":["%s","%s"]}`,
		personName, personAge, personAtoms, personNicknames[0], personNicknames[1])
}

func TestJsonDecoding(t *testing.T) {
	cx := makeTestContextWithContentTypeAndBody("application/json", makeJsonData())
	// check the "context" param is available (incase it consumes body)
	if cx.GetRequestContext() != "123" {
		t.Errorf("GetRequestContext() should return the correct request context before cx.Fill")
	}
	// fill struct 
	person := new(personTestStruct)
	err := cx.Fill(person)
	if err != nil {
		t.Errorf("json-decoder:", err)
	}
	// check it
	if person.Name != personName {
		t.Errorf("json-decoder: expected %v got %v", personName, person.Name)
	}
	if person.Age != personAge {
		t.Errorf("json-decoder: expected %v got %v", personAge, person.Age)
	}
	if person.Atoms != personAtoms {
		t.Errorf("json-decoders: expected %v got %v", personAtoms, person.Atoms)
	}
	if !reflect.DeepEqual(person.Nicknames, personNicknames) {
		t.Errorf("json-decoders: expected %v got %v", personNicknames, person.Nicknames)
	}
	// check the "context" param is still available
	if cx.GetRequestContext() != "123" {
		t.Errorf("GetRequestContext() should return the correct request context after cx.Fill")
	}
}

// If we want to allow Fill to be called multiple times
// then we will need to store the request Body in a buffer
// func TestDoubleJsonFill(t *testing.T) {
// 	// tests if it is possible to decode the body multiple times
// 	cx := makeTestContextWithContentTypeAndBody("application/json", makeJsonData())
// 	// fill struct multiple times
// 	for i := 0; i < 2; i++ {
// 		person := new(personTestStruct)
// 		err := cx.Fill(person)
// 		if err != nil {
// 			t.Errorf("form-decoder:", err)
// 		}
// 		if person.Name != "Alice" {
// 			t.Errorf("form-decoder: expected 'alice' got %v", person.Name)
// 		}
// 	}
// 
// }

func makeXmlData() string {
	return fmt.Sprintf(`<Person>
        <name>%s</name>
        <age>%d</age>
        <atoms>%d</atoms>
        <Nicknames>%s</Nicknames>
        <Nicknames>%s</Nicknames>
    </Person>`, personName, personAge, personAtoms, personNicknames[0], personNicknames[1])
}

func TestXmlDecoding(t *testing.T) {
	cx := makeTestContextWithContentTypeAndBody("application/xml", makeXmlData())
	// check the "context" param is available (incase it consumes body)
	if cx.GetRequestContext() != "123" {
		t.Errorf("GetRequestContext() should return the correct request context before cx.Fill")
	}
	// fill struct 
	var person personTestStruct
	err := cx.Fill(&person)
	if err != nil {
		t.Errorf("xml-decoder:", err)
	}
	// check it
	if person.Name != personName {
		t.Errorf("xml-decoder: expected %v got %v", personName, person.Name)
	}
	if person.Age != personAge {
		t.Errorf("xml-decoder: expected %v got %v", personAge, person.Age)
	}
	if person.Atoms != personAtoms {
		t.Errorf("xml-decoders: expected %v got %v", personAtoms, person.Atoms)
	}
	if !reflect.DeepEqual(person.Nicknames, personNicknames) {
		t.Errorf("xml-decoders: expected %v got %v", personNicknames, person.Nicknames)
	}
	// check the "context" param is still available
	if cx.GetRequestContext() != "123" {
		t.Errorf("GetRequestContext() should return the correct request context after cx.Fill")
	}
}

func TestUnknownDecoding(t *testing.T) {
	cx := makeTestContextWithContentTypeAndBody("application/junk", "<<junk>>")
	// fill struct 
	person := new(personTestStruct)
	err := cx.Fill(person)
	if err == nil {
		t.Errorf("form-decoder: should have raised error for unknown content-type: application/junk")
	}
}
