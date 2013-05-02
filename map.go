package goweb

/*
  Map maps an executor to the specified PathPattern on the DefaultHttpHandler.
*/
func Map(options ...interface{}) error {
	return DefaultHttpHandler().Map(options...)
}

/*
  DEVNOTE: This function is not tested because it simply passes the call on to the
  DefaultHttpHandler.
*/
