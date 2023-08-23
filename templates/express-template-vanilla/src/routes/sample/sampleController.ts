import { Request, Response } from "express";
import { getSamplesFunc } from "../../services/sampleService";

export function makeGetSamplesRoute(getSamples: getSamplesFunc) {
    return async function getSamplesRoute(req: Request, res: Response) {
        try {
            const samplesResponse = await getSamples();
            res.status(200).json(samplesResponse);
        } catch (err) {
            // check error logic
            res.status(500).json({message: (err as Error).message})
        }
    };
}

