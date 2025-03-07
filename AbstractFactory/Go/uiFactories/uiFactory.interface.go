package uiFactories

import (
	"abstractFactory/uiComponents/buttons"
	"abstractFactory/uiComponents/checkboxes"
	"errors"
)

type IUiFactory interface {
	GetButton() buttons.IButton
	GetCheckbox() checkboxes.ICheckbox
}

func GetUiFactory(os string) (IUiFactory, error) {
	if os == "MacOs" {
		return &MacUiFactory{}, nil
	} else if os == "Windows" {
		return &WindowsUiFactory{}, nil
	} else {
		return nil, errors.New("os not supported")
	}
}
