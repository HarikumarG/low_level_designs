import { DuckWrapper } from "./duckWrapper";

export class Main {
    run(): void {
        const duckWrapper = new DuckWrapper;
        duckWrapper.run(3);
    }
}
