package main

import "abstractFactory/uiFactories"

func main() {
	macFactory, _ := uiFactories.GetUiFactory("MacOs")
	macButton := macFactory.GetButton()
	macCheckbox := macFactory.GetCheckbox()

	macButton.RenderButton()
	macCheckbox.RenderCheckbox()
	macCheckbox.CloseCheckbox()
	macButton.CloseButton()

	windowsFactory, _ := uiFactories.GetUiFactory("Windows")
	windowsButton := windowsFactory.GetButton()
	windowsCheckbox := windowsFactory.GetCheckbox()

	windowsButton.RenderButton()
	windowsCheckbox.RenderCheckbox()
	windowsCheckbox.CloseCheckbox()
	windowsButton.CloseButton()
}
