import { FlyBehaviour } from "./types";

export class NoFly implements FlyBehaviour {
    fly(): String {
        return "No fly method from fly behaviour";
    };
};
