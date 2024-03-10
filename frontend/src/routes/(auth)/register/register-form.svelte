<script lang="ts">
    import { Label } from "$lib/components/ui/label";
    import type { ActionData } from "./$types";
    import { Input } from "$lib/components/ui/input";
    import { Button } from "$lib/components/ui/button";
    import Loader2 from "lucide-svelte/icons/loader-2";
    // import SuperDebug from "sveltekit-superforms";
    import {
        registerFormSchema,
        type RegisterFormSchema,
    } from "./register-form";
    import { type SuperValidated, superForm } from "sveltekit-superforms";
    import { zodClient } from "sveltekit-superforms/adapters";
    import { toast } from "svelte-sonner";

    export let data: {
        form: SuperValidated<RegisterFormSchema>;
    };
    const registerForm = superForm(data.form, {
        validators: zodClient(registerFormSchema),
        multipleSubmits: "prevent",
        delayMs: 500,
    });
    export let form: ActionData;

    const { form: formData, errors, enhance, delayed } = registerForm;
</script>

<form method="POST" action="?/register" use:enhance>
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
        {#if form?.status === 409}
            <p class="text-red-600 test-sm">User already exists</p>
        {/if}
        {#if !$delayed}
            <Button on:click|once type="submit" class="w-full"
                >Create account</Button
            >
        {:else if $delayed}
            <Button disabled>
                <Loader2 class="mr-2 h-4 w-4 animate-spin" />
                Creating account...
            </Button>
        {/if}
    </div>
</form>

<!-- <SuperDebug data={formData} /> -->
