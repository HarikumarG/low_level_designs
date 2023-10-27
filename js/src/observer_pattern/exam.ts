import { Observable } from "./observable";
import { Observer } from "./observer";
import { Subject } from "./subject";
import { ExamTime } from "./types";

export class Exam implements Observable {
    
    private observers: Array<Observer> = [];
    private examTime: ExamTime;
    
    add(observer: Observer): void {
        this.observers.push(observer)
    }

    notify(): void {
        this.observers.forEach((observer) => {
            observer.updates(this.examTime)
        })
    }

    setExamTime(subject: Subject, time: String) {
        this.examTime = { subject, time };
    }
}
