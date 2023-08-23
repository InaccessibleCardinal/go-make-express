import { Request, Response } from 'express';
import { makeGetSamplesRoute } from './sampleController';

jest.mock('express', () => {
    return {
        NextFunction: jest.fn(),
        Request: {},
        Response: {},
    };
});

function makeMockResponse() {
    const mockRes = {} as Response;
    mockRes.status = jest.fn(() => mockRes);
    mockRes.json = jest.fn();
    return mockRes;
}

const mockSamples = [{
    name: "sampleOne",
    attr1: "attribute one",
    attr2: "attribute two",
    subSample: { problems: 99, answer: 42 }
}];

beforeEach(() => {
    jest.resetAllMocks();
});

describe("makeGetSamplesRoute function", () => {
    it("should inject getSamples and return a route handler, happy path", async () => {
        const mockReq = {} as Request;
        const mockRes = makeMockResponse();
        const getSamplesSvcFunc = async () => {
            return mockSamples;
        };
        const handler = makeGetSamplesRoute(getSamplesSvcFunc);
        await handler(mockReq, mockRes);
        expect(mockRes.status).toHaveBeenCalledWith(200);
        expect(mockRes.json).toHaveBeenCalledWith(mockSamples);
    });

    it("should inject getSamples and return a route handler, failure path", async () => {
        const mockErr = new Error("lol")
        const mockReq = {} as Request;
        const mockRes = makeMockResponse();
        const getSamplesSvcFunc = async () => {
            throw mockErr;
        };
        const handler = makeGetSamplesRoute(getSamplesSvcFunc);
        await handler(mockReq, mockRes);

        expect(mockRes.status).toHaveBeenCalledWith(500);
    });
});