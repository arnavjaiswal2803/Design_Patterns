interface Observable {
    register(observer : Observer) : void;
    unregister(observer : Observer) : void;
    notify() : void;
}

class IphoneObservable implements Observable {
    private data : number;
    private  observers : Observer[];

    constructor() {
        this.observers = [];
    }

    public register(observer: Observer): void {
        this.observers.push(observer);
    }

    public unregister(observer: Observer): void {
        const index  = this.observers.indexOf(observer);
        this.observers.splice(index, 1);
    }

    public notify(): void {
        for (let i = 0; i < this.observers.length; i++) {
            this.observers[i].update(this);
        }
    }

    public setData(data : number){
        this.data = data;
        this.notify();
    }

    public getData() {
        return this.data;
    }
}

interface Observer {
    update(observable : Observable) : void;
}

class EmailObserver implements Observer {
    private username : string;

    constructor(username: string) {
        this.username = username;
    }

    public update(observable: Observable): void {
        if (observable instanceof IphoneObservable) {
            this.sendEmail("iphone");
        }
    }

    private sendEmail(item : string) {
        console.log(`sending email to ${this.username} for ${item}.`);
    }
    
}

class SmsObserver implements Observer {
    private phoneNumber : string;

    constructor(phoneNumber : string) {
        this.phoneNumber = phoneNumber;
    }

    public update(observable: Observable): void {
        if (observable instanceof IphoneObservable) {
            this.sendSms("iphone");
        }
    }

    private sendSms(item : string) {
        console.log(`sending sms to ${this.phoneNumber} for ${item}.`);
    }
    
}

const emailObserver = new EmailObserver("abcd@gmail.com");
const smsObserver = new SmsObserver("999xxxxx99");

const iphoneObservable = new IphoneObservable();
iphoneObservable.register(emailObserver);
iphoneObservable.register(smsObserver);

iphoneObservable.setData(5);

iphoneObservable.unregister(emailObserver);

iphoneObservable.setData(10);