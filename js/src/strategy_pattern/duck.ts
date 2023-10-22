import { FlyBehaviour } from "./types";

export class Duck {
    quack(): String {
        return "Duck quack method";
    };
    
    display(): String {
        return "Duck display method";
    };

    fly(flyObj: FlyBehaviour): String {
        return flyObj.fly();
    };
};
