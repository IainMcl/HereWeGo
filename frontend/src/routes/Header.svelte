<script lang="ts">
      import * as Avatar from "$lib/components/ui/avatar";
      import * as Command from "$lib/components/ui/command";
</script>

<header class="bg-pink-500">
      <Command.Root>
            <Command.Input placeholder="Type a command or search..." />
            <Command.List>
                  <Command.Empty>No results found.</Command.Empty>
                  <Command.Group heading="Suggestions">
                        <Command.Item>Calendar</Command.Item>
                        <Command.Item>Search Emoji</Command.Item>
                        <Command.Item>Calculator</Command.Item>
                  </Command.Group>
                  <Command.Separator />
                  <Command.Group heading="Settings">
                        <Command.Item>Profile</Command.Item>
                        <Command.Item>Billing</Command.Item>
                        <Command.Item>Settings</Command.Item>
                  </Command.Group>
            </Command.List>
      </Command.Root>
      <Avatar.Root>
            <Avatar.Image src="https://github.com/shadcn.png" alt="Shadcn" />
            <Avatar.Fallback>Avtar</Avatar.Fallback>
      </Avatar.Root>
</header>
