import { Bird } from "./bird";

export class Crow implements Bird {
    speak(): void {
        console.log("Kraaa kraa");
    }
};
