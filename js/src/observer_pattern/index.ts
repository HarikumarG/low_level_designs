import { Exam } from "./exam";
import { Student } from "./student";
import { Subject } from "./subject";

export class Main {
    run() {
        const student1 = new Student("Hari")
        const student2 = new Student("Ravi");

        const exam = new Exam
        exam.add(student1);
        exam.add(student2);

        const subject1 = new Subject("Science");
        exam.setExamTime(subject1, "2023-10-26");
        exam.notify();

        const subject2 = new Subject("Social");
        exam.setExamTime(subject2, "2023-10-30");
        exam.notify();
    }
}
