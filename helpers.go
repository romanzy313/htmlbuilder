package htmlbuilder

import (
	"context"
	"fmt"
	"strings"
)

func RenderElementToString(elem Renderable) (string, error) {
	var sb strings.Builder
	err := elem.Render(context.Background(), &sb)
	if err != nil {
		return "", fmt.Errorf("error rendering element: %v", err)
	}
	return sb.String(), nil
}
