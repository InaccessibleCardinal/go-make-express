import { Request, Response } from "express";
import { getSamplesFunc } from "../../services/sampleService";

export function makeGetSamplesRoute(getSamples: getSamplesFunc) {
    return async function getSamplesRoute(req: Request, res: Response) {
        const samplesResponse = await getSamples();
        return samplesResponse
            .map(value => res.status(200).json(value))
            .mapErr(err => {
                // check error logic
                res.status(500).json({message: err.message});
            });
    };
}

