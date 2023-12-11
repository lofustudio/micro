<script lang="ts">
    import Sidebar from "$lib/components/sidebar.svelte";
    import { ModeWatcher } from "mode-watcher";

    import "../app.css";
</script>

<ModeWatcher />
<div
    class="flex flex-row w-full min-h-screen background bg-white/75 dark:bg-stone-950"
>
    <Sidebar />
    <slot />
</div>
