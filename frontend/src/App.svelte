<script>
    import SideLeading from "./icons/SideLeading.svelte";
    import { cubicInOut } from 'svelte/easing';
    import { tweened } from 'svelte/motion';
    import { onMount } from "svelte";

    export let sidebarWidth = 300;
    export let sidebarMinWidth = 200;
    export let sidebarMaxWidth;
    let draging = false;
    const motionSidebarWidth = tweened(sidebarWidth, {
        duration: 200,
        easing: cubicInOut,
    });

    $: editorWidth = $motionSidebarWidth;
    $: if ($motionSidebarWidth <= 110) {
        editorWidth = 110;
    }

    onMount(() => {
        document.addEventListener('mouseup', function() {
            draging = false;
            document.removeEventListener('mousemove', mousemoving);
        });
    });

    /**
     * @param {{ pageX: any; stopPropagation: () => void; preventDefault: () => void; }} e
     */
    function mousemoving(e) {
        let width = e.pageX;
        if (width <= sidebarMinWidth) {
            return;
        }
        if ((sidebarMaxWidth && width >= sidebarMaxWidth) || (width >= document.body.clientWidth / 2)) {
            return;
        }
        draging = true;
        e.stopPropagation();
        e.preventDefault();
        motionSidebarWidth.set(width, {duration: 0});
    }

    function toggleSidebar() {
        if ($motionSidebarWidth === 0) {
            motionSidebarWidth.set(sidebarWidth);
        } else {
            sidebarWidth = $motionSidebarWidth;
            motionSidebarWidth.set(0);
        }
    }
</script>

<div class="w-screen h-screen flex select-none">
    <div id="sidebar" class="h-full" style="width: {$motionSidebarWidth}px">
        <div id="sidetool" class="h-[39px] min-w-[110px] absolute flex items-center pl-[81px] z-50" style="width: {$motionSidebarWidth}px">
            <SideLeading on:click={toggleSidebar} />
        </div>
        <div class="h-full w-full dark:text-white overflow-hidden flex flex-col" style="--wails-draggable:no-drag">
            <div class="flex-none h-[40px] border-b-[1px] border-solid border-gray-400 border-opacity-40 dark:border-opacity-20"></div>
            <div class="overflow-auto h-full">
                <div class="p-2">
                    hahah
                </div>
            </div>
        </div>
    </div>
    <div id="splitter" on:mousedown={() => document.addEventListener('mousemove', mousemoving)} class="splitter splitter-horizontal cursor-col-resize w-[1px] bg-gray-400 bg-opacity-40 dark:bg-black dark:bg-opacity-40" style="--wails-draggable:no-drag"/>
    <div class="grow h-full">
        <div style="width: calc(100% - {editorWidth}px)" id="toolbar"
            class="h-[39px] pl-[10px] pr-[11px] flex items-center justify-between absolute right-0 top-0 z-50">
            <strong id="title" class:select-none={draging} class="dark:text-white select-text" style="--wails-draggable:no-drag">Editor Header</strong>
            <SideLeading />
        </div>
        <div id="content" class:select-none={draging} class="mt-[39px] h-full w-full border-t-[1px] border-solid border-gray-400 border-opacity-40 dark:border-opacity-20 dark:text-white select-text" style="--wails-draggable:no-drag">
            Content View
        </div>
    </div>
</div>

<style lang="scss">
    .select-none {
        user-select: none;
    }

    .splitter {
        position: relative;
        &:before {
            content: "";
            position: absolute;
            left: 0;
            top: 0;
            opacity: 0;
            z-index: 1;
        }
        &:hover:before {
            opacity: 1;
        }

        &.splitter-horizontal:before {
            left: -5px;
            right: -5px;
            height: 100%;
            cursor: row-resize;
        }
    }
</style>
