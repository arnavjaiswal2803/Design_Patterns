package checkboxes

import "fmt"

type Checkbox struct {
	Os string
}

func (b *Checkbox) RenderCheckbox() {
	fmt.Printf("\nRendering Checkbox on %s", b.Os)
}

func (b *Checkbox) CloseCheckbox() {
	fmt.Printf("\nClosing Checkbox on %s", b.Os)
}
