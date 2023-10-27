import { Main as Strategy } from "./strategy_pattern";
import { Main as Observer } from "./observer_pattern";

console.log("Strategy Pattern");
new Strategy().run();
console.log("\n");

console.log("Observer Pattern");
new Observer().run();
console.log("\n");
