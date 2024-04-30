package htmlbuilder

import (
	"context"
	"strings"
)

type XSelectOptionMap map[string]*Element

type XSelectElement struct {
	*Element

	items         XSelectOptionMap
	placeholder   *Element
	selectedValue string
}

func XSelect(items XSelectOptionMap) *XSelectElement {
	return &XSelectElement{
		Element: NewElement("select"),
		items:   items,
	}
}

func (x *XSelectElement) SelectedValue(value string) *XSelectElement {
	x.selectedValue = value
	return x
}

func (x *XSelectElement) Default(defaultV *Element) *XSelectElement {
	defaultV.Attribute("disabled", "").Attribute("hidden", "")
	x.placeholder = defaultV
	return x
}

func (x *XSelectElement) Render(ctx context.Context, b *strings.Builder) error {

	if x.placeholder != nil {
		x.Element.Child(x.placeholder) // .Element is not needed here, but just for clarity
	}

	wasSelected := false
	// create all children to add to own selement
	for k, v := range x.items {
		if x.selectedValue == k {
			v.Attribute("selected", "")
			wasSelected = true
		}
		x.Element.Child(v)
	}

	if !wasSelected && x.placeholder != nil {
		x.placeholder.Attribute("selected", "")
	}

	return x.Element.Render(ctx, b)

}
