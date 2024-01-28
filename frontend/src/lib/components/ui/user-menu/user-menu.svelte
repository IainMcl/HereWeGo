<script lang="ts">
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
    import * as Avatar from "$lib/components/ui/avatar";
</script>

<DropdownMenu.Root>
    <DropdownMenu.Trigger>
        <Avatar.Root>
            <Avatar.Image src="https://github.com/shadcn.png" alt="Shadcn" />
            <Avatar.Fallback>AV</Avatar.Fallback>
        </Avatar.Root>
    </DropdownMenu.Trigger>
    <DropdownMenu.Content>
        <DropdownMenu.Item>Profile</DropdownMenu.Item>
        <DropdownMenu.Item>Billing</DropdownMenu.Item>
        <DropdownMenu.Item>Settings</DropdownMenu.Item>
    </DropdownMenu.Content>
</DropdownMenu.Root>
