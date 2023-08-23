import { NextFunction, Request, Response } from "express";

export function loggingMiddlewareSample(req: Request, res: Response, next: NextFunction) {
    console.info(`incoming request ${req.url}, ${req.method}`);
    next();
}

export function corsMiddlewareSample(req: Request, res: Response, next: NextFunction) {
    const corsHeaders = {
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Methods': 'GET, OPTIONS, POST, PUT, DELETE',
        'Access-Control-Allow-Headers': 'Authorization, Content-Type',
    };
    for (const [key, value] of Object.entries(corsHeaders)) {
        res.setHeader(key, value);
    }
    next();
} 