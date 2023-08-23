import "reflect-metadata";
import express from 'express';
import initializeRoutes from './routes';
import { createContainer } from "./createContainer";

const port = process.env.appPort || 3333;
const app = express();
const container = createContainer();

initializeRoutes(app, container);

app.listen(port, () => console.log(`app running at port ${port}`));