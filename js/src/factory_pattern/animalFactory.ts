import { Animal } from "./animal";

export interface AnimalFactory {
    createAnimal(x: number): Animal
};
