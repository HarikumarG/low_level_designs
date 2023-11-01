import { Animal } from "./animal";

export class Cat implements Animal {
    speak(): void {
        console.log("Meow");
    }
};
