import { Observer } from "./observer"

export interface Observable {
    add(s: Observer): void
    notify(): void
}
