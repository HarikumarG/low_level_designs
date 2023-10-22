import { Duck } from "./duck";
import { WildDuck } from "./wildDuck";
import { CityDuck } from "./cityDuck";
import { NoFly } from "./noFly";
import { WildCityFly } from "./wildCityFly";

export class DuckWrapper {
    run(x: number): void {
        let duck: Duck;
        if (x == 1) {
            duck = new CityDuck;
            console.log(duck.fly(new WildCityFly));
        } else if(x == 2) {
            duck = new WildDuck;
            console.log(duck.fly(new WildCityFly));
        } else {
            duck = new Duck();
            console.log(duck.fly(new NoFly));
        }
        console.log(duck.quack());
        console.log(duck.display());
    };
}
