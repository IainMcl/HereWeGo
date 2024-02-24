<script lang="ts">
    import { Label } from "$lib/components/ui/label";
    import { Input } from "$lib/components/ui/input";
    import { Button } from "$lib/components/ui/button";
    import SuperDebug from "sveltekit-superforms";
    import {
        registerFormSchema,
        type RegisterFormSchema,
    } from "./register-form";
    import {
        type SuperValidated,
        type Infer,
        superForm,
    } from "sveltekit-superforms";
    import { zodClient } from "sveltekit-superforms/adapters";

    export let data: {
        form: SuperValidated<RegisterFormSchema>;
    };
    const registerForm = superForm(data.form, {
        validators: zodClient(registerFormSchema),
        multipleSubmits: "prevent",
    });

    export let form: any;

    const { form: formData, errors, enhance } = registerForm;
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
            <p class="text-red-600">
                User already exists <a href="#">log in?</a>
            </p>
        {/if}
        <Button on:click|once type="submit" class="w-full"
            >Create account</Button
        >
    </div>
</form>

<!-- <SuperDebug data={formData} /> -->
