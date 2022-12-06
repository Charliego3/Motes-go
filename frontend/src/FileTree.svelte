<script>
    import IoIosArrowForward from 'svelte-icons/io/IoIosArrowForward.svelte'
    import { slide } from 'svelte/transition';

    export let name = null;
    export let path = name;
    export let isDir = false;
    export let createAt = "";
    export let updateAt = "";
    export let depth = 0;
    export let children;
    export let expanded = false;
    export let onClick = null; // click handler
    export let open = null;
    let selected = false;
    let padding = "";
    $: {
        if (depth > 0) {
            padding = "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;".repeat(depth-1);
        }
    }

    let disabled = true;
    let selectedNode = null;
    function toggle() {
        onClick(selectedNode, isDir);
    }

    function dbClick() {
        if (isDir && !!children) {
            expanded = !expanded;
            return;
        }

        open(name, path);
    }
</script>

{#if !!name}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div bind:this={selectedNode} class:selected={selected}
         on:dblclick={dbClick}
         class="h-[24px] overflow-hidden pr-3 pl-[6.5px] flex items-center"
         on:click={toggle}
         style="font-size: 80%;">
        {@html padding}
        {#if !!children}
            <div class="flex-none" style="width: 15px;height: 17px;animation-duration: 1s; animation-name: {expanded ? 'rotate-open' : 'rotate-close'};" 
                 on:click={dbClick}><IoIosArrowForward/></div>
        {:else}
            <div class="flex-none" style="width: 15px;"></div>
        {/if}
        {#if isDir}
            {#if !!children}
                <svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
                    width="17.2705" height="13.4985" viewBox="0 0 17 13.4985" class="ml-[3px] mr-[8px] flex-none">
                    <g>
                        <path d="M0 11.1328C0 12.6489 0.769043 13.4033 2.2998 13.4033L14.5532 13.4033C15.8643 13.4033 16.6333 12.6416 16.6333 11.1328L16.6333 5.12695L0 5.12695ZM0 4.16748L16.6333 4.16748L16.6333 3.61084C16.6333 2.10205 15.8569 1.34033 14.3335 1.34033L7.30225 1.34033C6.8042 1.34033 6.50391 1.22314 6.1377 0.908203L5.69092 0.541992C5.20752 0.131836 4.83398 0 4.10889 0L2.02148 0C0.739746 0 0 0.732422 0 2.20459Z" 
                            class="unselected-folder"/>
                    </g>
                </svg>
            {:else}
                <svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
                     width="16.6333" height="13.4985" viewBox="0 0 17 13.4985" class="ml-[3px] mr-[8px] flex-none">
                    <g>
                        <rect height="13.4985" opacity="0" width="16.6333" x="0" y="0"/>
                        <path d="M2.2998 13.4033L14.5532 13.4033C15.8643 13.4033 16.6333 12.6416 16.6333 11.1328L16.6333 3.61084C16.6333 2.10205 15.8569 1.34033 14.3335 1.34033L7.30225 1.34033C6.8042 1.34033 6.50391 1.22314 6.1377 0.908203L5.69092 0.541992C5.20752 0.131836 4.83398 0 4.10889 0L2.02148 0C0.739746 0 0 0.732422 0 2.20459L0 11.1328C0 12.6489 0.769043 13.4033 2.2998 13.4033ZM2.31445 12.2241C1.58203 12.2241 1.1792 11.8359 1.1792 11.0742L1.1792 2.27051C1.1792 1.54541 1.56006 1.17188 2.26318 1.17188L3.80859 1.17188C4.29199 1.17188 4.58496 1.29639 4.9585 1.61133L5.40527 1.98486C5.88135 2.38037 6.26953 2.51953 6.99463 2.51953L14.3115 2.51953C15.0366 2.51953 15.4541 2.91504 15.4541 3.67676L15.4541 11.0815C15.4541 11.8359 15.0366 12.2241 14.3115 12.2241ZM0.717773 5.2002L15.9082 5.2002L15.9082 4.09424L0.717773 4.09424Z" 
                        class="unselected-folder"/>
                    </g>
                </svg>
            {/if}
        {:else}
            <svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
                width="17.2705" height="13.4839" viewBox="0 0 17 13.4985" class="ml-[3px] mr-[8px] flex-none">
                <g>
                    <rect height="13.4839" opacity="0" width="17.2705" x="0" y="0"/>
                    <path d="M3.73535 3.69141L9.93164 3.69141C10.1953 3.69141 10.3857 3.49365 10.3857 3.2373C10.3857 2.98828 10.1953 2.79053 9.93164 2.79053L3.73535 2.79053C3.47168 2.79053 3.28125 2.98828 3.28125 3.2373C3.28125 3.49365 3.47168 3.69141 3.73535 3.69141ZM3.73535 10.686L9.93164 10.686C10.1953 10.686 10.3857 10.4956 10.3857 10.2466C10.3857 9.99023 10.1953 9.78516 9.93164 9.78516L3.73535 9.78516C3.47168 9.78516 3.28125 9.99023 3.28125 10.2466C3.28125 10.4956 3.47168 10.686 3.73535 10.686ZM2.2998 13.4839L14.9707 13.4839C16.5088 13.4839 17.2705 12.7295 17.2705 11.2207L17.2705 2.27051C17.2705 0.761719 16.5088 0 14.9707 0L2.2998 0C0.769043 0 0 0.761719 0 2.27051L0 11.2207C0 12.7295 0.769043 13.4839 2.2998 13.4839ZM2.31445 12.3047C1.58203 12.3047 1.1792 11.9165 1.1792 11.1548L1.1792 2.33643C1.1792 1.57471 1.58203 1.1792 2.31445 1.1792L14.9561 1.1792C15.6812 1.1792 16.0913 1.57471 16.0913 2.33643L16.0913 11.1548C16.0913 11.9165 15.6812 12.3047 14.9561 12.3047ZM3.0835 8.4668L14.1943 8.4668C14.8022 8.4668 15.1611 8.08594 15.1611 7.48535L15.1611 5.99121C15.1611 5.39062 14.8022 5.00977 14.1943 5.00977L3.0835 5.00977C2.47559 5.00977 2.1167 5.39062 2.1167 5.99121L2.1167 7.48535C2.1167 8.08594 2.47559 8.4668 3.0835 8.4668ZM3.74268 7.19238C3.479 7.19238 3.28125 6.99463 3.28125 6.73828C3.28125 6.48926 3.479 6.2915 3.74268 6.2915L7.34619 6.2915C7.60986 6.2915 7.80029 6.48926 7.80029 6.73828C7.80029 6.99463 7.60986 7.19238 7.34619 7.19238Z" 
                    class="unselected-folder"/>
                </g>
            </svg>
        {/if}
        <input value="{name}" {disabled} 
               class:text-black={!disabled} 
               class:bg-transparent={disabled} 
               class:ring-transparent={disabled}
               class:ring-offset-transparent={disabled}
               class:select-none={disabled}
               class="mr-2"
               style="text-overflow: ellipsis;width: 100%;line-height: normal;"/>
        <!-- {depth} -->
    </div>
{/if}

{#if !!children && expanded}
    <ul transition:slide>
        {#each children as node}
            <li>
                <svelte:self {open} {...node} children="{node.children}" depth="{++depth}" onClick={onClick}></svelte:self>
            </li>
        {/each}
    </ul>
{/if}

<style>
    ul {
        margin: 0;
        padding: 0;
    }

    li {
        padding: 0;
        margin: 0;
    }

    @keyframes rotate-open {
        from {
            transform: rotate(0);
        }
        to {
            transform: rotate(-90deg);
        }
    }

    @keyframes rotate-close {
        from {
            transform: rotate(-90deg);
        }
    }

    :global(.selected-folder) {
        fill: #FFFFFF;
    }

    @media (prefers-color-scheme: dark) {
        .unselected-folder {
            fill: #bcbcbc;
        }
    }

    @media (prefers-color-scheme: light) {
        .unselected-folder {
            fill: #525253;
        }
    }

    .selected {
        border-radius: 5px;
    }

    @media (prefers-color-scheme: dark) {
        .selected {
            background-color: #568DED;
        }
    }

    @media (prefers-color-scheme: light) {
        .selected {
            color: white;
            background-color: #1879E5;
            opacity: .9;
        }
    }
</style>
