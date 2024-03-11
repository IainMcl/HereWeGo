<script lang="ts">
    import { Toaster } from "$lib/components/ui/sonner";
</script>

<div class="split-fade grid place-content-center">
    <div
        class="z-0 bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500 h-screen w-screen object-cover bg-cover bg-center bg-no-repeat"
    ></div>
    <div class="z-10 absolute grid place-items-center h-screen w-screen">
        <slot />
    </div>
    <Toaster />
</div>
