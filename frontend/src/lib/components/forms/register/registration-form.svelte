<script lang="ts">
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    import { Button } from "$lib/components/ui/button";

    import { unauthenticatedRequest } from "@/services/api-service/apiService";

    const register = async (event: Event) => {
        event.preventDefault();
        const form = event.target as HTMLFormElement;
        const email = form.email.value;
        const password = form.password.value;
        console.log(email, password);

        const response = await unauthenticatedRequest(
            "/auth/register",
            "POST",
            {
                email,
                password,
            },
        );

        console.log(response);
    };
</script>

<section.Root>
    <h1.Title>Register</h1.Title>
    <form on:submit={register}>
        <Label for="email">Email</Label>
        <Input type="email" id="email" name="email" required />
        <Label for="password">Password</Label>
        <Input type="password" id="password" name="password" required />
        <Button type="submit">Register</Button>
    </form>
</section.Root>
