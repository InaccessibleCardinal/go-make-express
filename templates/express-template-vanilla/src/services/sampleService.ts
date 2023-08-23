import { Sample } from "../models/sample";

export interface getSamplesFunc {
    (): Promise<Sample[]>
}

export async function getSamples(): Promise<Sample[]> {
    if (process.env.makeErr) {
        throw new Error("failed to get samples");
    }
    const samples: Sample[] = [
        { name: "sampleOne", attr1: "attribute one", attr2: "attribute two", subSample: { problems: 99, answer: 42 } },
        { name: "sampleTwo", attr1: "attribute two one", attr2: "attribute two one", subSample: { problems: 99, answer: 42 } }
    ];
    return samples;
}