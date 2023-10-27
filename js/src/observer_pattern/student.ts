import { Observer } from "./observer";
import { ExamTime } from "./types";

export class Student implements Observer {
    
    name: String
    constructor(name: string) {
        this.name = name;
    }

    updates(examTime: ExamTime): void {
        console.log(`Hi ${this.name}, for subject ${examTime.subject.getName()} the exam time is ${examTime.time}`)
    }
}
