class Animal {
    constructor(name) {
        this.name = name;
    }
    speak() {
        console.log("hi ,i'm " + this.name);
    }
}
class dog extends Animal {
    constructor(name) {
        super(name);
        this.name = name;
    }
    bark() {
        console.log(this.name + "发出了吼叫Woof!");
    }
}
const p = new Animal('aaa');
const q = new dog("bbb");
p.speak();
q.speak();
q.bark();