import { Animal } from "./animal";
import { AnimalFactory } from "./animalFactory";
import { Cat } from "./cat";
import { Dog } from "./dog";

export class RandomAnimalFactory implements AnimalFactory {
    createAnimal(x: number): Animal {
        // some runtime operation/calculation
        if(x % 2 == 0) {
            return new Dog();
        }
        return new Cat();
    }
};
