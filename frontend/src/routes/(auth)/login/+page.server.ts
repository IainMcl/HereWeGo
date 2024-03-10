import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { fail } from '@sveltejs/kit';
import { loginFormSchema } from './login-form';
import { unauthenticatedRequest } from '$lib/services/api-service/apiService';
import type { PageServerLoad, Actions } from './$types';
// import type { Request } from '@sveltejs/kit';

export const load: PageServerLoad = async () => {
    return {
        status: 200,
    };
}

export const actions: Actions = {
    login: async ({ request }) => {
        const form = await superValidate(request, zod(loginFormSchema));
        if (!form.valid) {
            return fail(400, {
                form,
            });
        }

        const { email, password } = form.data;
        const response = await unauthenticatedRequest('POST', '/auth/login', {
            email,
            password,
        });

        return {
            form,
        };
    }
}