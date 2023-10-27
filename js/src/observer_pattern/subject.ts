export class Subject {
    private subjectName: String

    constructor(name: String) {
        this.subjectName = name;
    }

    getName() {
        return this.subjectName;
    }
}
