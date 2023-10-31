import { Calculation } from "./calculation";

export class FirstCalculation implements Calculation {
    getNumber(): number {
        return 10;
    }
};
