export interface Sample {
    name: string;
    attr1: string;
    attr2: string;
    subSample: SubSample;
}

export interface SubSample {
    problems: number;
    answer: number;
}
