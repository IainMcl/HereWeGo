import { z } from "zod";

const registerSchema = z.object({
    email: z.string().email({
        message: "Invalid email address",
    }),
    password: z.string().min(6, {
        message: "Password must be at least 6 characters long",
    }),
    confirmPassword: z.string().min(6, {
        message: "Password must be at least 6 characters long",
    }),
});