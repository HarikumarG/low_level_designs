import { Animal } from "./animal";
import { AnimalBirdFactory } from "./animalBirdFactory";
import { Bird } from "./bird";
import { Cat } from "./cat";
import { Crow } from "./crow";
import { Dog } from "./dog";
import { Eagle } from "./eagle";

export class RandomAnimalFactory implements AnimalBirdFactory {
    createAnimal(x: number): Animal {
        // some runtime operation/calculation
        if(x % 2 == 0) {
            return new Dog();
        }
        return new Cat();
    };
    createBird(x: number): Bird {
        // some runtime operation/calculation
        if(x % 2 == 0) {
            return new Eagle();
        }
        return new Crow();
    }
};
