This is a prototype of how an HTML builder in Go could be implemented. Inspired by various similar libraries, all trying to do the same thing: enable typesafe HTML templating without a build step.

Here is an example of how it could be used:

```go
func TestSmoke(t *testing.T) {
    div := Div().Class("container").Class("another").Text("yaar").Children(
        Div().Text("Hello, world!"),
        Textf("Hello, world %s", "again"),
    )
    res := div.String()

    expected := `<div class="container another">yaar<div>Hello, world!</div>Hello, world again</div>`

    if res != expected {
        t.Errorf("Expected %s, got %s", expected, res)
    }
}
```

Other libraries in no particular order:

- [htmlgo](https://github.com/julvo/htmlgo)
- [elem-go](https://github.com/chasefleming/elem-go)
- [go-app](https://github.com/maxence-charriere/go-app)
- [gohtml](https://github.com/yosssi/gohtml)
- [hb](https://github.com/gouniverse/hb)