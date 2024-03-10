import type { PageServerLoad, Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { registerFormSchema } from './register-form';
import { unauthenticatedRequest } from '@/services/api-service/apiService';
import { toast } from 'svelte-sonner';

export const load: PageServerLoad = async () => {
    return {
        form: await superValidate(zod(registerFormSchema)),
        status: 200,
    };
}

export const actions: Actions = {
    register: async ({ request }) => {
        const form = await superValidate(request, zod(registerFormSchema));
        if (!form.valid) {
            console.log("form not valid")
            return fail(400, {
                form,
            });
        }

        const { email, password } = form.data;
        const response = await unauthenticatedRequest('POST', '/auth/register', {
            email,
            password,
        });

        switch (response.status) {
            case 201:
                console.log("user created");
                let resp = await response.json();
                let createdEmail = resp.email;
                toast.success('User created', { description: `User ${createdEmail} has been created. You can now login.` });
                // Redirect to login page /login
                redirect(303, "/login")
            case 409:
                console.log("user already exists")
                return fail(409, {
                    status: 409,
                    body: {
                        error: 'User already exists',
                    },
                    form
                });
            default:
                return fail(500, {
                    status: 500,
                    body: {
                        error: 'Internal server error',
                    },
                    form
                });
        };
    },
};