/*
References:
    https://github.com/kamranahmedse/design-patterns-for-humans#behavioral-design-patterns
    https://youtube.com/playlist?list=PLrhzvIcii6GNjpARdnO4ueTUAVR9eMBpc&si=OlP2f1BUmzR3ojP_
*/
import { Main as Strategy } from "./strategy_pattern";
import { Main as Observer } from "./observer_pattern";
import { Main as Decorator } from "./decorator_pattern";
import { Main as Factory } from "./factory_pattern";
import { Main as AbstractFactory } from "./abstract_factory_pattern";

console.log("Strategy Pattern");
new Strategy().run();
console.log("\n");

console.log("Observer Pattern");
new Observer().run();
console.log("\n");

console.log("Decorator Pattern");
new Decorator().run();
console.log("\n");

console.log("Factory Pattern");
new Factory().run();
console.log("\n");

console.log("Abstract Factory Pattern");
new AbstractFactory().run();
console.log("\n");
