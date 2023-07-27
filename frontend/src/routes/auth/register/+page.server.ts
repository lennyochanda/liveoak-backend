import { fail, redirect } from '@sveltejs/kit'
import { BASE_API_URI } from '$lib/utils'
import type { Actions, PageServerLoad, } from '../$types';

// export const load: PageServerLoad = ({ locals }) => {
//     if (locals.user) {
//         throw redirect(302, '/')
//     }
// }

export const actions: Actions = {
    register: async ({ cookies, request, locals }) => {
        const data = Object.fromEntries(await request.formData());

        if (!data.email || !data.password || !data.username) {
            return fail(400, {
                error: "Missing email and or password"
            })
        }

        const { email, password, username } = data as { email: string, password: string, username: string };

        const requestInitOptions = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: username,
                email: email,
                password: password
            })
        }

        const res = await fetch(`${BASE_API_URI}/v1/user/create`, requestInitOptions)

        if (!res.ok) {
            return fail(500, {
                error: "An error occured"
            })
        }

        throw redirect(302, '/auth/login')
    }
}