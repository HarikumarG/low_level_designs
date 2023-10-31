import { AddDecorator } from "./addDecorator";
import { SubDecorator } from "./subDecorator";
import { FirstCalculation } from "./firstCalculation";

export class Main {
    run() {
        const firstCalc = new FirstCalculation;
        console.log(firstCalc.getNumber());

        const addCalc = new AddDecorator(firstCalc, 2);
        console.log(addCalc.getNumber());

        const subCalc = new SubDecorator(addCalc, 5);
        console.log(subCalc.getNumber());
    }
}
