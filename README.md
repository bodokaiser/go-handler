# Handler

Collection of common [http.Handler](http://godoc.org/net/http#Handler) for
[go](http://golang.org).

## Usage

### Auth

`Auth` wraps a `http.Handler` and responds `Unauthorized` to all requests which
do not match HTTP Basic authentication.

    a := &Auth{
        Username: "foo",
        Password: "bar",
        Handler:  mymux,
    }
    http.Handle("/", a)

### Logger

`Logger` wraps a `http.Handler` and prints the requests method and url.

    l := &Logger{
        Handler: mymux,
    }
    http.Handle("/", l)

### NotFound

`NotFound` responds `Not Found` using
[go-context](http://github.com/satisfeet/go-context) style responses.

Find more details by reading the
[godoc](http://godoc.org/github.com/satisfeet/go-handler).

## License

Copyright 2014 Bodo Kaiser <i@bodokaiser.io>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
