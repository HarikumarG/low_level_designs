import { Animal } from "./animal";
import { RandomAnimalFactory } from "./randomAnimalFactory";

export class Main {
    run() {
        const factory = new RandomAnimalFactory();
        const dog: Animal = factory.createAnimal(2);
        dog.speak();
        const cat: Animal = factory.createAnimal(3);
        cat.speak();
    };
};
