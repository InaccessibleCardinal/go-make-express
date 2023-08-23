import { Router } from 'express';
import { makeGetSamplesRoute } from './sampleController';
import { getSamples } from '../../services/sampleService';
import { loggingMiddlewareSample, corsMiddlewareSample } from '../../middleware';

const sampleRouter = Router();

sampleRouter.get("/",  loggingMiddlewareSample, corsMiddlewareSample, makeGetSamplesRoute(getSamples));

export default sampleRouter;

