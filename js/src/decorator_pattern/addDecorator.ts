import { Calculation } from "./calculation";

export class AddDecorator implements Calculation {

    calculation: Calculation;
    value: number;

    constructor(c: Calculation, v: number) {
        this.calculation = c;
        this.value = v;
    }

    getNumber(): number {
        return this.calculation.getNumber() + this.value;
    }
};
