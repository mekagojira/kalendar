package defined

type Context[I any, O any] struct {
	Id         string
	Input      *I
	Output     *O
	Code       string
	StatusCode int
	Message    string
}

func (c *Context[I, O]) WithOutput(output *O) *Context[I, O] {
	c.Output = output
	c.Code = "OK"
	c.StatusCode = 200
	return c
}

func (c *Context[I, O]) WithError(code string, message string, statusCode ...int) *Context[I, O] {
	c.Code = code
	c.Message = message

	if len(statusCode) > 0 {
		c.StatusCode = statusCode[0]
	} else {
		c.StatusCode = 400
	}

	return c
}

func (c *Context[I, O]) BadRequest() *Context[I, O] {
	c.Code = "BAD_REQUEST"
	c.Message = "Bad Request"
	return c
}

func (c *Context[I, O]) InternalServerError() *Context[I, O] {
	c.Code = "INTERNAL_SERVER_ERROR"
	c.Message = "Internal Server Error"
	c.StatusCode = 500
	return c
}

func (c *Context[I, O]) Unauthorized() *Context[I, O] {
	c.Code = "UNAUTHORIZED"
	c.Message = "Unauthorized"
	c.StatusCode = 401
	return c
}
