import { Application,Request,  Response } from 'express';
import { Routes } from './routes';
import sampleRouter from './sample';

export default function initializeRoutes(app: Application) {
    app.use(Routes.samples, sampleRouter);
    app.use(Routes.root, (req: Request, res: Response) => res.json({message: "welcome to the app"}));
}