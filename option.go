package htmlbuilder

type OptionElement struct {
	*Element
}

func Option(value string) *OptionElement {
	e := NewElement("option")

	e.Attribute("value", value)

	return &OptionElement{
		Element: e,
	}
}

// issue OptionElement is lost...
func (o *OptionElement) Selected() *OptionElement {
	o.Element.Attribute("selected", "")
	return o
}

func (o *OptionElement) SelectedIf(cond bool) *OptionElement {
	if cond {
		o.Element.Attribute("selected", "")
	}
	return o
}

func (o *OptionElement) SelectedIfFunc(cond func() bool) *OptionElement {
	if cond() {
		o.Element.Attribute("selected", "")
	}
	return o
}
