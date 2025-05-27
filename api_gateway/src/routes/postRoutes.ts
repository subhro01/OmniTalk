import { createProxyMiddleware, Options } from 'http-proxy-middleware';
import express from 'express';
import { IncomingMessage, ServerResponse } from 'http';


const router = express.Router();
const POST_SERVICE_URL = process.env.POST_SERVICE_URL || 'http://post-service:8001';

const proxyOptions: Options = {
    target: POST_SERVICE_URL,
    changeOrigin: true
}

router.use('/posts', createProxyMiddleware(proxyOptions))

export default router;