import type { PageServerLoad } from "$types";
import { z } from "zod";
import { zod } from "sveltekit-superforms/adapters";
import { superValidate } from "sveltekit-superforms";

const loginSchema = z.object({
    email: z.string().email({
        message: "Invalid email address",
    }),
    password: z.string().min(6, {
        message: "Password must be at least 6 characters long",
    }),
});

export default loginSchema;

export const load: PageServerLoad = (async () => {
    const form = await superValidate(zod(loginSchema));
    return { form };
});