import express from 'express';
import axios from 'axios';

const router = express.Router();
const AUTH_SERVICE_URL = process.env.AUTH_SERVICE_URL || 'http://localhost:8000';

router.post('/register', async (req, res) => {
    try {
        const response = await axios.post(`${AUTH_SERVICE_URL}/api/auth/register/`, req.body);
        res.status(response.status).json(response.data);
    } catch (err: any) {
        res.status(err.response?.status || 500).json(err.response?.data || {error: "Auth service error"});
    }
})

router.post('/login', async (req, res) => {
    try {
        const response = await axios.post(`${AUTH_SERVICE_URL}/api/auth/login/`, req.body);
        res.status(response.status).json(response.data);
    } catch (err: any) {
        res.status(err.response?.status || 500).json(err.response?.data || {error: "Auth service error"});
    }
})

router.post('/verify-token', async (req, res) => {
    try {
        const response = await axios.post(`${AUTH_SERVICE_URL}/api/auth/verify-token/`, req.body);
        res.status(response.status).json(response.data);
    } catch (err: any) {
        res.status(err.response?.status || 500).json(err.response?.data || {error: "Auth service error"});
    }
})

export default router;