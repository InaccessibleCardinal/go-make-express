import 'reflect-metadata';
import { Request, Response } from 'express';
import { SampleController } from './SampleController';
import { ISampleService } from '../../services/ISampleService';
import { Sample } from '../../models/sample';

const mockSamples = [{
    name: "sampleOne",
    attr1: "attribute one",
    attr2: "attribute two",
    subSample: { problems: 99, answer: 42 }
}];

const mockErr = new Error("lol");

jest.mock('express', () => {
    return {
        NextFunction: jest.fn(),
        Request: {},
        Response: {},
    };
});

class TestSampleServiceSuccess implements ISampleService {
    public async getSamples(): Promise<Sample[]> {
        return mockSamples;
    }
}

class TestSampleServiceFailure implements ISampleService {
    public async getSamples(): Promise<Sample[]> {
        throw mockErr;
    }
}

function makeMockResponse() {
    const mockRes = {} as Response;
    mockRes.status = jest.fn(() => mockRes);
    mockRes.json = jest.fn();
    return mockRes;
}

beforeEach(() => {
    jest.resetAllMocks();
});

describe("makeGetSamplesRoute function", () => {
    it("should inject getSamples and return a route handler, happy path", async () => {
        const mockReq = {} as Request;
        const mockRes = makeMockResponse();
        const testSampleController = new SampleController(new TestSampleServiceSuccess());
        await testSampleController.getAll(mockReq, mockRes);
        expect(mockRes.status).toHaveBeenCalledWith(200);
        expect(mockRes.json).toHaveBeenCalledWith(mockSamples);
    });

    it("should inject getSamples and return a route handler, failure path", async () => {
        const mockReq = {} as Request;
        const mockRes = makeMockResponse();
        const testSampleController = new SampleController(new TestSampleServiceFailure());
        await testSampleController.getAll(mockReq, mockRes);
        expect(mockRes.status).toHaveBeenCalledWith(500);
    });
});