package htmlbuilder

import (
	"context"
	"strings"
	"testing"
)

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

func TestWriterImplementation(t *testing.T) {

	div := Div().Class("container").Class("another").Text("yaar").Children(
		Div().Text("Hello, world!"),
		Textf("Hello, world %s", "again"),
	)

	res := div.StringWriter()
	// t.Logf("Result: %s", res)

	expected := `<div class="container another">yaar<div>Hello, world!</div>Hello, world again</div>`

	// t.Errorf("Expected <div><div>Hello, world!</div></div>, got %s", res)

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestStringRendering(t *testing.T) {

	div := DoctypeHTML(
		Div().Text("Hello, world!"),
	)

	res := div.String()

	expected := `<!DOCTYPE html><div>Hello, world!</div>`

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestRunner(t *testing.T) {
	tests := []struct {
		name     string
		element  Renderable
		expected string
	}{
		{
			name: "SmokyDoky",
			element: Div().Class("container").Class("another").Text("yaar").Children(
				Div().Text("Hello, world!"),
				Textf("Hello, world %s", "again"),
			),
			expected: `<div class="container another">yaar<div>Hello, world!</div>Hello, world again</div>`,
		},
		{
			name: "Option",
			element: Fragment().Children(
				Option("1").SelectedIf(true).Element.Id("test").Text("One"),
				Option("2").Element.Id("test").Text("Two"),
			),
			expected: `<option value="1" selected="" id="test">One</option><option value="2" id="test">Two</option>`,
		},
		{
			name:     "Input",
			element:  Input().Type("text").Name("name").Value("value").Placeholder("enter name").Id("name").Class("form-control"),
			expected: `<input type="text" name="name" value="value" placeholder="enter name" id="name" class="form-control"></input>`,
		},
		{
			name:     "BsButton",
			element:  BsButton().Primary().Small().Type("submit").OnClick(`alert('clicked')`).Id("running-it"),
			expected: `<button type="submit" onclick="alert('clicked')" id="running-it" class="btn btn-primary btn-sm"></button>`,
		},
		{
			name: "BsButtonStyled",
			element: BsButton().Styled(BsButtonStyle{
				Outline: "primary",
				Size:    "sm",
			}).Type("submit"),
			expected: `<button type="submit" class="btn btn-outline-primary btn-sm"></button>`,
		},
		{
			name:     "Doctype",
			element:  DoctypeHTML(Div().Text("Hello, world!")),
			expected: `<!DOCTYPE html><div>Hello, world!</div>`,
		},
	}

	for _, tt := range tests {

		b := strings.Builder{}

		tt.element.Render(context.Background(), &b)

		res := b.String()

		if res != tt.expected {
			t.Errorf("Test %s failed. Expected %s, got %s", tt.name, tt.expected, res)
		}
	}
}
