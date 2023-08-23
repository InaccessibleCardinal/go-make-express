import { Request, Response } from "express";
import {injectable, inject} from "tsyringe";
import { ISampleService } from "../../services/ISampleService";

@injectable()
export class SampleController {
    constructor(@inject("ISampleService") private service: ISampleService) {}
    public getAll = async (_req: Request, res: Response) => {
        try {
            const samples = await this.service.getSamples();
            res.status(200).json(samples);
        } catch (err) {
            // check error logic
            console.error(err);
            res.status(500).json({message: (err as Error).message});
        }
    }
}