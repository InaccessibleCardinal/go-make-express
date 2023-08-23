import {container} from "tsyringe";
import { SampleService } from "./services/SampleService";
import { SampleController } from "./routes/sample/SampleController";

export function createContainer() {
    container.register('ISampleService', {useClass: SampleService});
    container.register('SampleController', {useClass: SampleController});
    return container;
}