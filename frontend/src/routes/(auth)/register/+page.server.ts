import type { PageServerLoad, Actions } from './$types';
import { fail } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { registerFormSchema } from './register-form';
import { unauthenticatedRequest } from '@/services/api-service/apiService';

export const load: PageServerLoad = async () => {
    return {
        form: await superValidate(zod(registerFormSchema))
    };
}

export const actions: Actions = {
    register: async ({ request }) => {
        console.log("register action called")
        // console.log(request);
        // const form = await request.formData();
        const form = await superValidate(request, zod(registerFormSchema));
        // console.log(form);
        if (!form.valid) {
            console.log("form not valid")
            return fail(400, {
                form,
            });
        }

        // let email = form.get('email');
        // let password = form.get('password');
        const { email, password } = form.data;
        const response = await unauthenticatedRequest('POST', '/auth/register', {
            email,
            password,
        });

        console.log(response);
        switch (response.status) {
            case 201:
                return {
                    status: 201,
                    headers: {
                        // 'set-cookie': response.headers.get('set-cookie'),
                    },
                    form
                };
            case 409:
                return {
                    status: 409,
                    body: {
                        error: 'User already exists',
                    },
                    form
                };
        }
        return {
            form,
        };
    },
};