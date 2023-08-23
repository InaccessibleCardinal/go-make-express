import express, {Request, Response} from 'express';

const port = process.env.appPort || 3333;
const app = express();

app.use('/samples', handleSamples);

app.listen(port, () => console.log(`app running at port ${port}`));

function handleSamples(req: Request, res: Response) {
    res.status(200).json([
        { name: "sampleOne", attr1: "attribute one", attr2: "attribute two", subSample: { problems: 99, answer: 42 } },
        { name: "sampleTwo", attr1: "attribute two one", attr2: "attribute two one", subSample: { problems: 99, answer: 42 } }
    ]);
}