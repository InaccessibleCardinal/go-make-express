import express from 'express';
import initializeRoutes from './routes';

const port = process.env.appPort || 3333;
const app = express();

initializeRoutes(app);

app.listen(port, () => console.log(`app running at port ${port}`));