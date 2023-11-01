import { Animal } from "./animal";
import { Bird } from "./bird";

export interface AnimalBirdFactory {
    createAnimal(x: number): Animal
    createBird(x: number): Bird
};
