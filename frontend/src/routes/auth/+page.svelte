<script lang="ts">
    import * as Form from "$lib/components/ui/form";
    import { registerSchema, type RegisterSchema } from "./schema";
    import type { SuperValidated } from "sveltekit-superforms";

    export let form: SuperValidated<RegisterSchema>;
</script>

<section>
    <Form.Root method="POST" {form} schema={registerSchema} let:config>
        <Form.Field {config} name="email">
            <Form.Item>
                <Form.Label>Email</Form.Label>
                <Form.Input />
                <Form.Description>Registation email address</Form.Description>
                <Form.Validation />
            </Form.Item>
        </Form.Field>
        <Form.Field {config} name="password">
            <Form.Item>
                <Form.Label>Password</Form.Label>
                <Form.Input />
                <Form.Description>Enter your password</Form.Description>
                <Form.Validation />
            </Form.Item>
        </Form.Field>
        <Form.Button>Submit</Form.Button>
    </Form.Root>
</section>
