import { Animal } from "./animal";
import { Bird } from "./bird";
import { RandomAnimalFactory } from "./randomAnimalFactory";

export class Main {
    run() {
        const factory = new RandomAnimalFactory();
        const dog: Animal = factory.createAnimal(2);
        dog.speak();
        const cat: Animal = factory.createAnimal(3);
        cat.speak();

        const eagle: Bird = factory.createBird(2);
        eagle.speak();
        const crow: Bird = factory.createBird(3);
        crow.speak();
    };
};
