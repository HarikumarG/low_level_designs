import { Bird } from "./bird";

export class Eagle implements Bird {
    speak(): void {
        console.log("Keeeee");
    };
};
