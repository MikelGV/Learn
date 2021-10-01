// External imports
import express from "express";
import bodyParser from "body-parser";

// My imports
import feedRoutes from "./routes/feed";


const app = express();

app.use(feedRoutes)


app.listen(3000);