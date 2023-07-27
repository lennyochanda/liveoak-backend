<script lang="ts">
    import { enhance } from '$app/forms';
    // import type {ActionData} from './$types'

    // export let form: ActionData;
</script>

<h1>
    Live Oak Login
</h1>

<form method="post" action="?/login" use:enhance>
    <label for="email">Email</label>
    <input type="email" name="email" id="email">
    <br>
    <label for="password">Password</label>
    <input type="password" name="password" id="password">
    <br>
    <button>Submit</button>

    <!-- {#if form?.error}
        <div>
            {form.error}
        </div>
    {/if} -->
</form>
