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
    login: async ({ cookies, url, request }) => {
        const form = await superValidate(request, zod(loginFormSchema));
        if (!form.valid) {
            return fail(400, {
                form,
            });
        }

        const { email, password } = form.data;
        const { headers, status } = await login(email, password);

        const setCookie = headers.get('set-cookie');
        const [cookie, path, httpOnly] = setCookie.split('; ');

        const [name, value] = cookie.split('=');
        const [_, pathValue] = path.split('=');
        const isHttpOnly = httpOnly === 'HttpOnly';

        cookies.set(name, value, {
            path: pathValue,
            httpOnly: isHttpOnly
        });

        switch (status) {
            case 200:
                const redirectTo = url.searchParams.get("redirectTo");
                if (redirectTo) {
                    return {
                        status: 302,
                        headers: {
                            location: `/${redirectTo.slice(1)}`,
                            'set-cookie': setCookie
                        },
                    };
                }
                return {
                    status: 302,
                    headers: {
                        location: '/',
                        'set-cookie': setCookie
                    },
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