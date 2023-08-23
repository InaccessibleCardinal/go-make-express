import { Application,Request,  Response } from 'express';
import { DependencyContainer } from 'tsyringe';
import { Routes } from './routes';
import {makeSampleRouter} from './sample';

export default function initializeRoutes(app: Application, container: DependencyContainer) {
    const sampleRouter = makeSampleRouter(app, container);
    app.use(Routes.samples, sampleRouter);
    app.use(Routes.root, (req: Request, res: Response) => res.json({message: "welcome to the app"}));
}