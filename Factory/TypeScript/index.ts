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

abstract class ButtonFactory {
    abstract getButton() : Button;
};

class WindowsButtonFactory extends ButtonFactory{
    getButton(): Button {
        return new WindowsButton();
    }
};

class MacButtonFactory extends ButtonFactory{
    getButton(): Button {
        return new MacButton();
    }
};

const winButtonFactory = new WindowsButtonFactory();
const windowsButton = winButtonFactory.getButton();

const macButtonFactory = new MacButtonFactory();
const macButton = macButtonFactory.getButton();


macButton.render();
windowsButton.render();
windowsButton.close();
macButton.close();