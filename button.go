package htmlbuilder

// button

type ButtonElement struct {
	*Element
}

func Button() *ButtonElement {
	return &ButtonElement{
		Element: NewElement("button"),
	}
}

// type can use a nice little const
func (b *ButtonElement) Type(typ string) *ButtonElement {
	b.Attribute("type", typ)
	return b
}
