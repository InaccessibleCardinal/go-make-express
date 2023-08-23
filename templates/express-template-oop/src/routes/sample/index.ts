import { Application, Router } from 'express';
import { DependencyContainer } from 'tsyringe';
import { SampleController } from './SampleController';

export function makeSampleRouter(app: Application, container: DependencyContainer) {
    const sampleRouter = Router();
    const sampleController = container.resolve('SampleController');
    sampleRouter.get("/", (sampleController as SampleController).getAll);
    return sampleRouter;
}


