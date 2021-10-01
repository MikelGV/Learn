import path from 'path';
import express from 'express';

import { Router } from 'express';

const router = Router();

router.get('/', (req, res ,next) => {
    res.send('Hello')
})

export default router