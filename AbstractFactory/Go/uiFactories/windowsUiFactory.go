package uiFactories

import (
	"abstractFactory/uiComponents/buttons"
	"abstractFactory/uiComponents/checkboxes"
)

type WindowsUiFactory struct {
}

func (f *WindowsUiFactory) GetButton() buttons.IButton {
	return &buttons.WindowsButton{
		Button: buttons.Button{
			Os: "Windows",
		},
	}
}

func (f *WindowsUiFactory) GetCheckbox() checkboxes.ICheckbox {
	return &checkboxes.MacCheckbox{
		Checkbox: checkboxes.Checkbox{
			Os: "Windows",
		},
	}
}
