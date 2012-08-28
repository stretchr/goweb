package goweb

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestAddSignedCookie(t *testing.T) {

	context := MakeTestContext()
	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = "2468"
	cookie.Path = "/something"
	cookie.Domain = "domain"
	cookie.RawExpires = "NOW"
	cookie.Expires = time.Now()
	cookie.MaxAge = 123
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.Raw = "userId=2468;"

	signedCookie, err := context.AddSignedCookie(cookie)

	if err != nil {
		t.Errorf("AddSignedCookie shouldn't return an error: %s", err)
		return
	}

	assertEqual(t, signedCookie.Name, fmt.Sprintf("%s_signed", cookie.Name), "Cookie name")
	assertEqual(t, signedCookie.Value, Hash(cookie.Value), "Cookie value (signed)")

	// assert the rest of the values were also copied
	assertEqual(t, signedCookie.Path, cookie.Path, "Path")
	assertEqual(t, signedCookie.Domain, cookie.Domain, "Domain")
	assertEqual(t, signedCookie.RawExpires, cookie.RawExpires, "RawExpires")
	assertEqual(t, signedCookie.Expires, cookie.Expires, "Expires")
	assertEqual(t, signedCookie.MaxAge, cookie.MaxAge, "MaxAge")
	assertEqual(t, signedCookie.Secure, cookie.Secure, "Secure")
	assertEqual(t, signedCookie.HttpOnly, cookie.HttpOnly, "HttpOnly")
	assertEqual(t, signedCookie.Raw, cookie.Raw, "Raw")

}

func TestSignedCookie(t *testing.T) {

	context := MakeTestContext()

	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = "2468"

	signedCookie, err := context.AddSignedCookie(cookie)

	if err != nil {
		t.Errorf("Shouldn't error: %s", err)
	}

	// set the request headers
	context.Request.Header = make(http.Header)
	context.Request.AddCookie(cookie)
	context.Request.AddCookie(signedCookie)

	returnedCookie, cookieErr := context.SignedCookie(cookie.Name)

	if cookieErr != nil {
		t.Errorf("SignedCookie shouldn't return error: %s", cookieErr)
		return
	}
	if returnedCookie == nil {
		t.Errorf("SignedCookie shouldn't return nil")
		return
	}

	assertEqual(t, returnedCookie.Name, cookie.Name, "name")

}

func TestSignedCookie_Tempered(t *testing.T) {

	context := MakeTestContext()

	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = "2468"

	signedCookie, err := context.AddSignedCookie(cookie)

	if err != nil {
		t.Errorf("Shouldn't error: %s", err)
	}

	// temper with the cookie
	cookie.Value = "something-else"

	// set the request headers
	context.Request.Header = make(http.Header)
	context.Request.AddCookie(cookie)
	context.Request.AddCookie(signedCookie)

	returnedCookie, cookieErr := context.SignedCookie(cookie.Name)

	if cookieErr == nil {
		t.Errorf("SignedCookie SHOULD return error")
	}
	if returnedCookie != nil {
		t.Errorf("ReturnedCookie should be nil")
	}

}

func TestValidateSignedCookie_Success(t *testing.T) {

	context := MakeTestContext()

	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = "2468"

	signedCookie, err := context.AddSignedCookie(cookie)

	if err != nil {
		t.Errorf("Shouldn't error: %s", err)
	}

	// set the request headers
	context.Request.Header = make(http.Header)
	context.Request.AddCookie(cookie)
	context.Request.AddCookie(signedCookie)

	valid, validErr := context.cookieIsValid(cookie.Name)

	if validErr != nil {
		t.Errorf("Shouldn't error: %s", validErr)
	}

	assertEqual(t, valid, true, "cookieIsValid should be true")

}

func TestValidateSignedCookie_Tampered(t *testing.T) {

	context := MakeTestContext()

	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = "2468"

	signedCookie, err := context.AddSignedCookie(cookie)

	if err != nil {
		t.Errorf("Shouldn't error: %s", err)
	}

	// tamper with the cookie value
	cookie.Value = "1357"

	// set the request headers
	context.Request.Header = make(http.Header)
	context.Request.AddCookie(cookie)
	context.Request.AddCookie(signedCookie)

	valid, validErr := context.cookieIsValid(cookie.Name)

	if validErr != nil {
		t.Errorf("Shouldn't error: %s", validErr)
	}

	assertEqual(t, valid, false, "cookieIsValid should be false")

}

func TestValidateSignedCookie_MissingCookie(t *testing.T) {

	context := MakeTestContext()

	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = "2468"

	signedCookie, err := context.AddSignedCookie(cookie)

	if err != nil {
		t.Errorf("Shouldn't error: %s", err)
	}

	// set the request headers
	context.Request.Header = make(http.Header)
	context.Request.AddCookie(signedCookie)

	valid, validErr := context.cookieIsValid(cookie.Name)

	if validErr != CookieIsMissing {
		t.Errorf("Should error CookieIsMissing")
	}

	assertEqual(t, valid, false, "cookieIsValid should be false")

}

func TestValidateSignedCookie_MissingSignedCookie(t *testing.T) {

	context := MakeTestContext()

	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = "2468"

	// set the request headers
	context.Request.Header = make(http.Header)
	context.Request.AddCookie(cookie)

	valid, validErr := context.cookieIsValid(cookie.Name)

	if validErr != SignedCookieIsMissing {
		t.Errorf("Should error SignedCookieIsMissing")
	}

	assertEqual(t, valid, false, "cookieIsValid should be false")

}
