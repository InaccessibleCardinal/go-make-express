import { Sample } from "../models/sample";

export interface ISampleService {
    getSamples(): Promise<Sample[]>
}