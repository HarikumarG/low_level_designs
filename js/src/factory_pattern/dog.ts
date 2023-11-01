import { Animal } from "./animal";

export class Dog implements Animal {
    speak(): void {
        console.log("Bark");
    }
};
