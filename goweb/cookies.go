package goweb

import (
	"errors"
	"fmt"
	"net/http"
)

var CookieIsMissing error = errors.New("Cookie is missing")
var SignedCookieIsMissing error = errors.New("Signed cookie is missing")

const (
	SignedCookieFormat string = "%s_signed"
)

func toSignedCookieName(name string) string {
	return fmt.Sprintf(SignedCookieFormat, name)
}

// AddSignedCookie adds the specified cookie to the response and also adds an
// additional 'signed' cookie that can be used to validate the cookies value later.
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
