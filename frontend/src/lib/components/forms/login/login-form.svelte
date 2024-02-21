<script lang="ts">
    import * as Form from "$lib/components/ui/form";
    import { superForm } from "sveltekit-superforms";
    import { load } from "./schema";

    const { form } = await load();
</script>

<Form.Root {form}></Form.Root>
