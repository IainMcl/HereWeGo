<script lang="ts">
    // import { Icons } from "$lib/components/icons";
    import type { PageData, ActionData } from "./$types";
    import RegisterForm from "./register-form.svelte";
    import * as Card from "$lib/components/ui/card";
    import { Button } from "$lib/components/ui/button";

    export let data: PageData;
    export let form: ActionData;
</script>

<svelte:head>
    <title>Register</title>
    <meta name="description" content="Register for an account" />
</svelte:head>

<Card.Root>
    <Card.Header class="space-y-1">
        <Card.Title class="text-2xl">Create an account</Card.Title>
        <Card.Description
            >Enter your email below to create your account</Card.Description
        >
    </Card.Header>
    <Card.Content class="grid gap-4">
        <!-- <div class="grid grid-cols-2 gap-6">
            <Button variant="outline">
                <Icons.gitHub class="mr-2 h-4 w-4" />
                GitHub
            </Button>
            <Button variant="outline">
                <Icons.google class="mr-2 h-4 w-4" /> 
                Google
            </Button>
        </div> -->
        <div class="relative">
            <div class="absolute inset-0 flex items-center">
                <span class="w-full border-t" />
            </div>
            <!-- <div class="relative flex justify-center text-xs uppercase">
                <span class="bg-card px-2 text-muted-foreground">
                    Or continue with
                </span>
            </div> -->
        </div>
        <RegisterForm {data} {form} />
        <div class="actions">
            <p class="text-center text-md">
                Already have an account?
                <Button
                    variant="link"
                    href="/login"
                    class="text-primary-500 underline">Login</Button
                >
            </p>
        </div>
    </Card.Content>
</Card.Root>
