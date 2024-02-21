import { z } from 'zod';

export const registerFormSchema = z.object({
    email: z.string().email({
        message: 'Please enter a valid email address',
    }),
    password: z.string().min(8, {
        message: 'Password must be at least 8 characters long',
    }),
});

export type RegisterFormSchema = z.infer<typeof registerFormSchema>;