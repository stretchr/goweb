package goweb

import (
	"errors"
	"fmt"
	"net/http"
)

// CookieIsMissing is returned when a cookie is missing.
var CookieIsMissing error = errors.New("Cookie is missing")

// SignedCookieIsMissing is returned when the signed cookie is missing.
var SignedCookieIsMissing error = errors.New("Signed cookie is missing")

// CookieNotValid is returned when the cookie and its signed counterpart do not match.
//
// I.e. the cookie value has been tampered with.
var CookieNotValid error = errors.New("Cookie and signed cookie do not match")

var (
	// SignedCookieFormat is the format string used to decide the name of the
	// signed cookie.
	SignedCookieFormat string = "%s_signed"
)

// toSignedCookieName gets the signed cookie name from the specified cookie name,
// by running it through the SignedCookieFormat string.
func toSignedCookieName(name string) string {
	return fmt.Sprintf(SignedCookieFormat, name)
}

// AddSignedCookie adds the specified cookie to the response and also adds an
// additional 'signed' cookie that is used to validate the cookies value when
// SignedCookie is called.
func (c *Context) AddSignedCookie(cookie *http.Cookie) (*http.Cookie, error) {

	// make the signed cookie
	signedCookie := new(http.Cookie)

	// copy the cookie settings
	signedCookie.Path = cookie.Path
	signedCookie.Domain = cookie.Domain
	signedCookie.RawExpires = cookie.RawExpires
	signedCookie.Expires = cookie.Expires
	signedCookie.MaxAge = cookie.MaxAge
	signedCookie.Secure = cookie.Secure
	signedCookie.HttpOnly = cookie.HttpOnly
	signedCookie.Raw = cookie.Raw

	// set the signed cookie specifics
	signedCookie.Name = toSignedCookieName(cookie.Name)
	signedCookie.Value = Hash(cookie.Value)

	// add the cookies
	http.SetCookie(c.ResponseWriter, cookie)
	http.SetCookie(c.ResponseWriter, signedCookie)

	// return the new signed cookie (and no error)
	return signedCookie, nil

}

// Gets the cookie specified by name and validates that its value has not been
// tampered with by checking the signed cookie too.  Will return CookieNotValid error
// if it has been tampered with, otherwise, it will return the actual cookie.
func (c *Context) SignedCookie(name string) (*http.Cookie, error) {

	valid, validErr := c.cookieIsValid(name)
	if valid {
		return c.Request.Cookie(name)
	} else if validErr != nil {
		return nil, validErr
	}

	return nil, CookieNotValid
}

// cookieIsValid checks to see if the cookie and its signed counterpart
// match.
func (c *Context) cookieIsValid(name string) (bool, error) {

	// get the cookies
	cookie, cookieErr := c.Request.Cookie(name)
	signedCookie, signedCookieErr := c.Request.Cookie(toSignedCookieName(name))

	// handle errors reading cookies
	if cookieErr == http.ErrNoCookie {
		return false, CookieIsMissing
	}
	if cookieErr != nil {
		return false, cookieErr
	}
	if signedCookieErr == http.ErrNoCookie {
		return false, SignedCookieIsMissing
	}
	if signedCookieErr != nil {
		return false, signedCookieErr
	}

	// check the cookies
	if Hash(cookie.Value) != signedCookie.Value {
		return false, nil
	}

	// success - therefore valid
	return true, nil

}
