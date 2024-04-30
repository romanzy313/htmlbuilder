package htmlbuilder

import (
	"testing"
)

// some complex elements cannot be rendered directly for whatever reason with String()
func TestXSelectFragmented(t *testing.T) {

	// cannot call .String() on this directly like so:
	// elem := XSelect(XSelectOptionMap{
	// 	"1": Option("1").Text("One"),
	// 	"2": Option("2").Text("Two"),
	// }).SelectedValue("2").Default(Option("").InnerHTML("Please select"))

	elem := Fragment(XSelect(XSelectOptionMap{
		"1": Option("1").Text("One"),
		"2": Option("2").Text("Two"),
	}).SelectedValue("2").Default(Option("").InnerHTML("Please select")))

	res := elem.String() // String cannot be called directly on XSelectElement, as its part of the Element... so need to wrap in a fragment

	expected := `<select><option value="" disabled="" hidden="">Please select</option><option value="1">One</option><option value="2" selected="">Two</option></select>`

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}

}

// ahh these are flakey tests, as the order of the options is not guaranteed
// it will look different on every render... yikes
// also this doesnt kelp when inspecting in dev tools
func TestXSelect(t *testing.T) {
	tests := []struct {
		name     string
		element  Renderable
		expected string
	}{
		{
			name:     "Empty items",
			element:  XSelect(XSelectOptionMap{}),
			expected: `<select></select>`,
		},
		{
			name:     "With attributes",
			element:  XSelect(XSelectOptionMap{}).Id("test"),
			expected: `<select id="test"></select>`,
		},
		{
			name: "Single item",
			element: XSelect(XSelectOptionMap{
				"1": Option("1").Text("One"),
			}),
			expected: `<select><option value="1">One</option></select>`,
		},
		{
			name: "With selected value",
			element: XSelect(XSelectOptionMap{
				"1": Option("1").Text("One"),
				"2": Option("2").Text("Two"),
			}).SelectedValue("2"),
			expected: `<select><option value="1">One</option><option value="2" selected="">Two</option></select>`,
		},
		{
			name: "With default value that is automatically selected",
			element: XSelect(XSelectOptionMap{
				"1": Option("1").Text("One"),
				"2": Option("2").Text("Two"),
			}).Default(Option("").InnerHTML("Please select")),
			expected: `<select><option value="" disabled="" hidden="" selected="">Please select</option><option value="1">One</option><option value="2">Two</option></select>`,
		},
		{
			name: "With default value where selected value is incorrect",
			element: XSelect(XSelectOptionMap{
				"1": Option("1").Text("One"),
				"2": Option("2").Text("Two"),
			}).Default(Option("").InnerHTML("Please select")).SelectedValue("3"),
			expected: `<select><option value="" disabled="" hidden="" selected="">Please select</option><option value="1">One</option><option value="2">Two</option></select>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, err := RenderElementToString(tt.element)

			if err != nil {
				t.Errorf("Test %s failed. Error: %v", tt.name, err)
			}

			if res != tt.expected {
				t.Errorf("Test %s failed. Expected \n%s \ngot \n%s", tt.name, tt.expected, res)
			}
		})
	}
}
