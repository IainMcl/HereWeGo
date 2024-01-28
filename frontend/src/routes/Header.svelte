<script lang="ts">
      import * as Command from "$lib/components/ui/command";
      import { UserMenu } from "$lib/components/ui/user-menu";
      import { ModeToggle } from "$lib/components/ui/mode-toggle";

      import { mode } from "mode-watcher";

      let logoSrc: string;
      $: logoSrc = $mode === "dark" ? "logoinv.png" : "logo.png";
</script>

<header
      class="sticky top-0 z-50 w-full border-b border-border/40 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60"
>
      <div class="container flex justify-between my-2">
            <div id="left">
                  <a href="/">
                        <img src={logoSrc} alt="Logo" class="h-8" />
                  </a>
            </div>
            <div id="header-right" class="flex gap-4 justify-end">
                  <!-- <Command.Root class="w-80">
                  <Command.Input placeholder="Type a command or search..." />
                  <Command.List>
                        <Command.Empty>No results found.</Command.Empty>
                        <Command.Group heading="Suggestions"></Command.Group>
                        <Command.Separator />
                        <Command.Group heading="Settings">
                              <Command.Item>Profile</Command.Item>
                              <Command.Item>Billing</Command.Item>
                              <Command.Item>Settings</Command.Item>
                        </Command.Group>
                  </Command.List>
            </Command.Root> -->
                  <ModeToggle />
                  <UserMenu />
            </div>
      </div>
</header>
