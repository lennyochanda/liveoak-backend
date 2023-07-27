import { fail, redirect } from '@sveltejs/kit'
import { BASE_API_URI } from '$lib/utils'
import type { Actions, PageServerLoad, } from '../$types';

// export const load: PageServerLoad = ({ locals }) => {
//     if (locals.user) {
//         throw redirect(302, '/')
//     }
// }

export const actions: Actions = {
    login: async ({ cookies, request, locals }) => {

        const data = Object.fromEntries(await request.formData());

        if (!data.email || !data.password) {
            return fail(400, {
                error: "Missing email and or password"
            })
        }

        const { email, password } = data as { email: string, password: string };

        const requestInitOptions = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email: email,
                password: password
            })
        }

        const res = await fetch(`${BASE_API_URI}/v1/user/login`, requestInitOptions)

        if (!res.ok) {
            return fail(500, {
                error: "An error occured"
            })
        }

        const {token} = await res.json();
        console.log(token);
        
        cookies.set('Authorization', `Bearer ${token}`, {
            httpOnly: true,
            sameSite: 'lax',
            secure: false,
            maxAge: 7200,
	    path: "/"
        })

        throw redirect(302, '/')
    },
}
