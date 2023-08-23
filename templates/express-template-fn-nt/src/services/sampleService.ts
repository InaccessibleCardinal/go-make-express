import { ok, Ok, err, Err } from 'neverthrow';
import { Sample } from "../models/sample";

export type OkSample = Ok<Sample[], never>;
export type ErrSample = Err<never, Error>;

export interface getSamplesFunc {
    (): Promise<OkSample | ErrSample>;
}

export async function getSamples(): Promise<OkSample | ErrSample> {
    if (process.env.makeErr) {
        return err(new Error("failed to get samples"));
    }
    const samples: Sample[] = [
        { name: "sampleOne", attr1: "attribute one", attr2: "attribute two", subSample: { problems: 99, answer: 42 } },
        { name: "sampleTwo", attr1: "attribute two one", attr2: "attribute two one", subSample: { problems: 99, answer: 42 } }
    ];
    return ok(samples);
}