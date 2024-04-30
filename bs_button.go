package htmlbuilder

type BsButtonElement struct {
	*ButtonElement
}

func BsButton() *BsButtonElement {
	btn := Button()
	btn.Class("btn")

	return &BsButtonElement{
		ButtonElement: btn,
	}
}

func (i *BsButtonElement) Kind(kind string) *BsButtonElement {
	i.Element.Classf("btn-%s", kind)
	return i
}

func (i *BsButtonElement) KindOutline(kind string) *BsButtonElement {
	i.Element.Classf("btn-outline-%s", kind)
	return i
}

func (i *BsButtonElement) Size(size string) *BsButtonElement {
	i.Element.Classf("btn-%s", size)
	return i
}

func (i *BsButtonElement) Large() *BsButtonElement {
	return i.Size("lg")
}
func (i *BsButtonElement) Small() *BsButtonElement {
	return i.Size("sm")
}

func (i *BsButtonElement) Primary() *BsButtonElement {
	i.Element.Class("btn-primary")
	return i
}

type BsButtonStyle struct {
	Kind    string
	Outline string
	Size    string
}

func (i *BsButtonElement) Styled(style BsButtonStyle) *BsButtonElement {
	if style.Outline != "" {
		i.KindOutline(style.Outline)
	}
	if style.Kind != "" {
		i.Kind(style.Kind)
	}
	if style.Size != "" {
		i.Kind(style.Size)
	}

	return i
}
