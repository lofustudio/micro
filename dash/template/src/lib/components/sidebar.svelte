<script>
    import ThemeSwitch from "./theme-switch.svelte";
    import * as Tooltip from "$lib/components/ui/tooltip";
</script>

<div
    class="flex flex-col p-4 gap-4 items-center border-r border-black/10 dark:border-white/10"
>
    <a href="/dash">
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="32"
            height="32"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            ><circle cx="12" cy="12" r="3" /><circle
                cx="19"
                cy="5"
                r="2"
            /><circle cx="5" cy="19" r="2" /><path
                d="M10.4 21.9a10 10 0 0 0 9.941-15.416"
            /><path d="M13.5 2.1a10 10 0 0 0-9.841 15.416" /></svg
        >
    </a>
    <div
        class="w-full px-3 border-b border-black/10 dark:border-white/10"
    ></div>
    <div class="flex flex-col h-full justify-between">
        <div class="flex flex-col justify-between gap-6 items-center">
            <Tooltip.Root openDelay={0}>
                <Tooltip.Trigger>
                    <a
                        href="/dash"
                        class="text-neutral-500 hover:text-white duration-300 transition-colors"
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="hover:text-white"
                            ><path d="M12 8V4H8" /><rect
                                width="16"
                                height="12"
                                x="4"
                                y="8"
                                rx="2"
                            /><path d="M2 14h2" /><path d="M20 14h2" /><path
                                d="M15 13v2"
                            /><path d="M9 13v2" /></svg
                        >
                    </a>
                </Tooltip.Trigger>
                <Tooltip.Content side="right" sideOffset={16}>
                    <span>Text to Speech</span>
                </Tooltip.Content>
            </Tooltip.Root>
        </div>
        <ThemeSwitch />
    </div>
</div>
