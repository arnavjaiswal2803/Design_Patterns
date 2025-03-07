package buttons

import "fmt"

type Button struct {
	Os string
}

func (b *Button) RenderButton() {
	fmt.Printf("\nRendering button on %s", b.Os)
}

func (b *Button) CloseButton() {
	fmt.Printf("\nClosing button on %s", b.Os)
}
