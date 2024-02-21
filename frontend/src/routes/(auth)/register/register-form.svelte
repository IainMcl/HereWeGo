<script lang="ts">
    import { Label } from "$lib/components/ui/label";
    import { Input } from "$lib/components/ui/input";
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

    export let data;
    console.log(data);
    const form = superForm(data, {
        validators: zodClient(registerFormSchema),
    });

    const { form: formData, errors, constraints, enhance } = form;
</script>

<form action="/register" method="post" use:enhance>
    <div class="form-sections flex gap-4 flex-col">
        <div class="form-input flex flex-col gap-2">
            <Label for="email">Email</Label>
            <!-- <Form.Label>Email</Form.Label> -->
            <Input type="email" bind:value={$formData.email} />
            {#if $errors.email}<span class="invalid text-red-600"
                    >{$errors.email}</span
                >{/if}
        </div>
        <div class="form-input flex flex-col gap-2">
            <Label for="password">Password</Label>
            <!-- <Form.Label>Email</Form.Label> -->
            <Input type="password" bind:value={$formData.Password} />
            {#if $errors.password}<span class="invalid text-red-600"
                    >{$errors.password}</span
                >{/if}
        </div>
    </div>
</form>
