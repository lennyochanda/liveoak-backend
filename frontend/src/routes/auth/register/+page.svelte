<script lang="ts">
    import { enhance } from '$app/forms';
    // import type {ActionData} from './$types'

    // export let form: ActionData;
</script>

<h2>LiveOak Register</h2>
<form method="post" action="?/register" use:enhance>
    <label for="username">Username</label>
    <input type="text" name="username" id="username">
    <br>
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