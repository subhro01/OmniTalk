import { Request, Response, NextFunction } from 'express';
import axios from 'axios';

const AUTH_SERVICE_URL=process.env.AUTH_SERVICE_URL

export const verifyToken = async (req: Request, res:Response, next:NextFunction) => {
    const token = req.headers['authorization'];

    if (!token) {
        return res.status(401).json({"error": "No token provided"});
    }

    try {
        const response = await axios.post(`${AUTH_SERVICE_URL}/api/auth/verify-token`, {
            token
        });
        (req as any).user = response.data.user;
    } catch (err: any) {
        res.status(err.response?.status || 500).json(err.response?.data || {error: "Token validation failed"});
    }
}