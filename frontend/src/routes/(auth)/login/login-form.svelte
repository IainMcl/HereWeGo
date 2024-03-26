<script lang="ts">
    import { Label } from "$lib/components/ui/label";
    import type { ActionData } from "./$types";
    import { Input } from "$lib/components/ui/input";
    import { Button } from "$lib/components/ui/button";
    import { loginFormSchema, type LoginFormSchema } from "./login-form";
    import { type SuperValidated, superForm } from "sveltekit-superforms";
    import { zodClient } from "sveltekit-superforms/adapters";
    import { toast } from "svelte-sonner";
    import { goto } from "$app/navigation";

    export let data: {
        form: SuperValidated<LoginFormSchema>;
    };
    const registerForm = superForm(data.form, {
        validators: zodClient(loginFormSchema),
        multipleSubmits: "prevent",
        delayMs: 500,
        onResult: ({ result }) => {
            console.log(result);
            switch (result.data.status) {
                case 400:
                    toast.error("Invalid email or password");
                    break;
                case 500:
                    toast.error("Server error");
                    break;
                case 302:
                    goto(result.data.headers.location);
                    break;
                default:
                    console.error(result);
                    toast.error("Unknown error");
            }
        },
    });
    // export let form: ActionData;

    const { form: formData, errors, enhance, delayed } = registerForm;
</script>

<form method="POST" action="?/login" use:enhance>
    <div class="form-sections flex gap-4 flex-col">
        <div class="form-input flex flex-col gap-2">
            <Label for="email">Email</Label>
            <Input name="email" type="email" bind:value={$formData.email} />
            {#if $errors.email}<span class="invalid text-red-600"
                    >{$errors.email}</span
                >{/if}
        </div>
        <div class="form-input flex flex-col gap-2">
            <Label for="password">Password</Label>
            <Input
                name="password"
                type="password"
                bind:value={$formData.password}
            />
            {#if $errors.password}<span class="invalid text-red-600"
                    >{$errors.password}</span
                >{/if}
        </div>
        {#if !$delayed}
            <Button on:click|once type="submit" class="w-full">Login</Button>
        {:else if $delayed}
            <Button disabled>Logging in...</Button>
        {/if}
    </div>
</form>

<!-- <SuperDebug data={formData} /> -->
