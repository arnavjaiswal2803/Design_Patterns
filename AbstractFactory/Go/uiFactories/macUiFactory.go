package uiFactories

import (
	"abstractFactory/uiComponents/buttons"
	"abstractFactory/uiComponents/checkboxes"
)

type MacUiFactory struct {
}

func (f *MacUiFactory) GetButton() buttons.IButton {
	return &buttons.MacButton{
		Button: buttons.Button{
			Os: "MacOs",
		},
	}
}

func (f *MacUiFactory) GetCheckbox() checkboxes.ICheckbox {
	return &checkboxes.MacCheckbox{
		Checkbox: checkboxes.Checkbox{
			Os: "MacOs",
		},
	}
}
