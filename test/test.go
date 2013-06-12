package test

func TestWeb(t *testing.T) {

  goweb.Test(t, "GET people/123", func(t *testing.T, response) {

    assert.Equal(t, 200, response.StatusCode())

  })

}

