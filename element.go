package htmlbuilder

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Renderable interface {
	Render(ctx context.Context, b *strings.Builder) error
	RenderWriter(w io.Writer) error
}

// this cant "get a type..."
type Element struct {
	Renderable
	name             string
	content          string
	classes          []string
	styles           []string
	attributes       []string
	childrenElements []Renderable

	// RenderSelf func() *Element // is this okay? children are render
}

func NewElement(name string) *Element {
	return &Element{
		name:             name,
		content:          "",
		classes:          make([]string, 0),
		styles:           make([]string, 0),
		attributes:       make([]string, 0),
		childrenElements: make([]Renderable, 0),
	}
}

func (e *Element) Child(child Renderable) *Element {
	e.childrenElements = append(e.childrenElements, child)
	return e
}

// render must render self and all children

func (e *Element) Children(children ...Renderable) *Element {
	e.childrenElements = append(e.childrenElements, children...)
	return e
}

func (e *Element) Id(id string) *Element {
	return e.Attribute("id", id)
}

func (e *Element) Class(class string) *Element {
	// this is given as a string, i dont need to keep kv here. I can convert to correct format and keep as string
	e.classes = append(e.classes, class) // just add a space after, ez
	return e
}

func (e *Element) Classf(format string, a ...any) *Element {
	e.classes = append(e.classes, fmt.Sprintf(format, a...))
	return e
}

// TODO: sanitize the key and value? maybe one maybe both
func (e *Element) Attribute(key, value string) *Element {
	// this is given as a string, i dont need to keep kv here. I can convert to correct format and keep as string
	e.attributes = append(e.attributes, fmt.Sprintf(`%s="%s"`, key, value))
	return e
}
func (e *Element) AttributeUnsafe(key, value string) *Element {
	// this is given as a string, i dont need to keep kv here. I can convert to correct format and keep as string
	e.attributes = append(e.attributes, fmt.Sprintf(`%s="%s"`, key, value))
	return e
}

// func (e *Element) RenderWriter(w io.Writer) {
// 	b := strings.Builder{}
// 	e.Render(&b)
// 	w.Write([]byte(b.String()))

// }

func (e *Element) RenderWriter(w io.Writer) error {
	isFragment := e.name == "" // fragment is empty

	if !isFragment {
		w.Write([]byte("<" + e.name))

		// first compile the class attribute

		if len(e.classes) > 0 {
			e.attributes = append(e.attributes, fmt.Sprintf(`class="%s"`, strings.Join(e.classes, " "))) // i dont like join, id like to write staight
		}

		if len(e.attributes) > 0 {
			w.Write([]byte(" "))
			for i, attr := range e.attributes {
				w.Write([]byte(attr))
				// only write non-last one
				if i < len(e.attributes)-1 {
					w.Write([]byte(" "))
				}
				// b.WriteString(" ")
			}
		}

		w.Write([]byte{'>'})
	}
	w.Write([]byte(e.content))
	for _, child := range e.childrenElements {
		// child.RenderRawer(b)
		err := child.RenderWriter(w)
		if err != nil {
			return err
		}
	}
	if !isFragment {
		// every single one of these could be an error....
		w.Write([]byte("</" + e.name + ">"))
	}

	return nil
}

func (e *Element) RenderToResponseWriter(w http.ResponseWriter, req *http.Request) error {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	ctx := req.Context()

	b := strings.Builder{}
	e.Render(ctx, &b)

	if ctx.Err() != nil {
		// handle the error
		return fmt.Errorf("context error: %v", ctx.Err())
	}

	// otherwise just render it
	_, err := w.Write([]byte(b.String()))

	if err != nil {
		// handle the error
		return err
	}

	return nil
}

func (e *Element) Render(ctx context.Context, b *strings.Builder) error {
	isFragment := e.name == "" // fragment is empty

	if !isFragment {

		b.WriteString("<" + e.name)

		// first compile the class attribute

		if len(e.classes) > 0 {
			e.attributes = append(e.attributes, fmt.Sprintf(`class="%s"`, strings.Join(e.classes, " "))) // i dont like join, id like to write staight
		}

		if len(e.attributes) > 0 {
			b.WriteString(" ")
			for i, attr := range e.attributes {
				b.WriteString(attr)
				// only write non-last one
				if i < len(e.attributes)-1 {
					b.WriteString(" ")
				}
				// b.WriteString(" ")
			}
		}

		b.WriteString(">")

	}
	b.WriteString(e.content)
	for _, child := range e.childrenElements {
		err := child.Render(ctx, b)
		if err != nil {
			return err
		}
	}
	if !isFragment {
		b.WriteString("</" + e.name + ">")
	}

	return nil
}

func (e *Element) StringWriter() string {
	var builder strings.Builder
	if err := e.RenderWriter(&builder); err != nil {
		// Handle the error appropriately. For simplicity, we'll return the error message,
		// but in a real application, you might want to handle it differently.
		return fmt.Sprintf("Error rendering HTML element: %v", err)
	}
	return builder.String()
}

// call it just String()
func (e *Element) String() string {
	b := strings.Builder{}

	err := e.Render(context.Background(), &b)
	if err != nil {
		return fmt.Sprintf("Error rendering HTML element: %v", err)
	}
	return b.String()
}
func (e *Element) StringWithContext(ctx context.Context) string {
	b := strings.Builder{}
	err := e.Render(ctx, &b)
	if err != nil {
		return fmt.Sprintf("Error rendering HTML element: %v", err)
	}
	return b.String()
}

// other nice to have methods

// TODO: sanitize me!
func (e *Element) Text(text string) *Element {
	// sanitize the value here!
	e.content = text
	return e
}

func (e *Element) Textf(format string, a ...any) *Element {
	e.content = fmt.Sprintf(format, a...)
	return e
}

func (e *Element) Unsafe(html string) *Element {
	e.content = html
	return e
}

func (e *Element) Unsafef(format string, a ...any) *Element {
	e.content = fmt.Sprintf(format, a...)
	return e
}

// This is an UNSAFE method, it will not sanitize the html
func (e *Element) InnerHTML(html string) *Element {
	e.content = html
	return e
}

// This is an UNSAFE method, it will not sanitize the html
func (e *Element) InnerHTMLf(format string, a ...any) *Element {
	e.content = fmt.Sprintf(format, a...)
	return e
}

// this is a special case, need to check it for sure.
// but the classes and shit dont need that...
func Fragment(children ...Renderable) *Element {
	e := &Element{
		name:             "",
		content:          "",
		classes:          nil,
		styles:           nil,
		attributes:       nil,
		childrenElements: nil,
	}

	if len(children) > 0 {
		e.Children(children...)
	}

	return e
}

// Text fragment helpers

func Text(text string) *Element {

	e := Fragment()
	// this also works, i just need to sanitize it
	// actually fragments can be so much more simplified, they dont need all that class and stuct initialization...
	e.Text(text)
	// e.Text(text)

	return e
}
func Textf(format string, a ...any) *Element {
	e := Fragment()
	e.Text(fmt.Sprintf(format, a...))

	return e
}

// unsafe fragment helpers

func Unsafe(rawhtml string) *Element {

	e := Fragment()
	// e.content = rawhtml
	// or
	e.InnerHTML(rawhtml)

	return e
}

func Unsafef(format string, a ...any) *Element {
	e := Fragment()
	e.content = fmt.Sprintf(format, a...)

	return e
}

// javascript

func (b *Element) OnClick(script string) *Element {
	// or can reference element directly
	b.Attribute("onclick", script)
	return b
}

// preamle

type DoctypeHTMLElement struct {
	*Element
}

func DoctypeHTML(body Renderable) *Element {
	return Fragment(
		Unsafe("<!DOCTYPE html>"),
		body,
	)
}

// DIV
type DivElement struct {
	*Element
}

// nothing special about this
func Div() *DivElement {
	return &DivElement{
		Element: NewElement("div"),
	}
}

// Input

type InputElement struct {
	*Element
}

func Input() *InputElement {
	return &InputElement{
		Element: NewElement("input"),
	}
}

func (i *InputElement) Type(typ string) *InputElement {
	i.Attribute("type", typ)
	return i
}

func (i *InputElement) Value(value string) *InputElement {
	i.Attribute("value", value)
	return i
}

func (i *InputElement) Placeholder(placeholder string) *InputElement {
	i.Attribute("placeholder", placeholder)
	return i
}

func (i *InputElement) Name(name string) *InputElement {
	i.Attribute("name", name)
	return i
}

// custom button (aka bootstrap styles)
