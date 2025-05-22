import express from 'express';
import dotenv from 'dotenv';
import authRoutes from './routes/authRoutes';


dotenv.config()
const app = express()
app.use(express.json())

app.use('/api/auth', authRoutes);

const PORT = process.env.PORT || 5000
app.listen(PORT, () => {
    console.log(`API Gateway running on ${PORT}`);
})