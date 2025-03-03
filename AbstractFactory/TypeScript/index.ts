interface Button {
    render(): void;
    close(): void;
};

class WindowsButton implements Button {
    render(): void {
        console.log("Rendering windows button");
    }

    close(): void {
        console.log("Closing windows button");
    }
};

class MacButton implements Button {
    render(): void {
        console.log("Rendering mac button");
    }

    close(): void {
        console.log("Closing mac button");
    }
};

interface Checkbox {
    render(): void;
    close(): void;
};

class WindowsCheckbox implements Checkbox {
    render(): void {
        console.log("Rendering windows Checkbox");
    }

    close(): void {
        console.log("Closing windows Checkbox");
    }
};

class MacCheckbox implements Checkbox {
    render(): void {
        console.log("Rendering mac Checkbox");
    }

    close(): void {
        console.log("Closing mac Checkbox");
    }
};

abstract class GUIFactory {
    abstract getButton() : Button;
    abstract getCheckbox() : Checkbox;
};

class WindowsGUIFactory extends GUIFactory{
    getButton(): Button {
        return new WindowsButton();
    }

    getCheckbox() : Checkbox {
        return new WindowsCheckbox();
    }
};

class MacGUIFactory extends GUIFactory{
    getButton(): Button {
        return new MacButton();
    }

    getCheckbox() : Checkbox {
        return new MacCheckbox();
    }
};

const winGUIFactory = new WindowsGUIFactory();
const windowsButton = winGUIFactory.getButton();
const windowsCheckbox = winGUIFactory.getCheckbox();

const macGUIFactory = new MacGUIFactory();
const macButton = macGUIFactory.getButton();
const macCheckbox = macGUIFactory.getCheckbox();


macButton.render();
windowsButton.render();
macCheckbox.render();
windowsCheckbox.render();

windowsCheckbox.close();
macCheckbox.close();
windowsButton.close();
macButton.close();