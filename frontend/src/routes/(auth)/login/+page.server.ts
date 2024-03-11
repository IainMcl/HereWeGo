import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { fail, redirect } from '@sveltejs/kit';
import { loginFormSchema } from './login-form';
import { login } from '$lib/services/api-service/apiService';
import type { PageServerLoad, Actions } from './$types';

export const load: PageServerLoad = async () => {
    return {
        form: await superValidate(zod(loginFormSchema)),
        status: 200,
    };
}

export const actions: Actions = {
    login: async ({ url, request }) => {
        const form = await superValidate(request, zod(loginFormSchema));
        if (!form.valid) {
            return fail(400, {
                form,
            });
        }

        const { email, password } = form.data;
        const { resp: response, token: token } = await login(email, password);

        switch (response.status) {
            case 200:
                const redirectTo = url.searchParams.get("redirectTo");
                if (redirectTo) {
                    return {
                        status: 302,
                        headers: {
                            location: `/${redirectTo.slice(1)}`,
                        },
                        token: token
                    };
                }
                return {
                    status: 302,
                    headers: {
                        location: '/',
                    },
                    token: token
                };
            case 401:
                return fail(401, {
                    form,
                });
            default:
                return {
                    form,
                };
        };
    }
}