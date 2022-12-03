<script>
    import SideLeading from "./icons/SideLeading.svelte";
    import { onMount } from "svelte";

    export let sidebarWidth = 300;
    export let sidebarMinWidth = 200;
    export let sidebarMaxWidth;

    onMount(() => {
        const splitter = document.getElementById('splitter');
        function dragable(width) {
            if (width <= sidebarMinWidth) {
                return false;
            }
            if ((sidebarMaxWidth && width >= sidebarMaxWidth) || (width >= document.body.clientWidth / 2)) {
                return false;
            }
            return true;
        }

        function mousemoving(e) {
            let width = e.pageX;
            if (!dragable(width)) {
                return
            }
            e.stopPropagation();
            e.preventDefault();
            const sidebar = document.getElementById("sidebar");
            sidebarWidth = width;
            sidebar.setAttribute('style', 'width: ' + width + 'px;');
        }

        splitter.addEventListener('mousedown', function(e) {
            document.getElementById('content').classList.replace('select-text', 'select-none');
            document.getElementById('title').classList.replace('select-text', 'select-none');
            document.addEventListener('mousemove', mousemoving);
        });
        document.addEventListener('mouseup', function() {
            document.getElementById('content').classList.replace('select-none', 'select-text');
            document.getElementById('title').classList.replace('select-none', 'select-text');
            document.removeEventListener('mousemove', mousemoving);
        });
    });

    function toggleSidebar() {
        const sidebar = document.getElementById("sidebar");
        const toolbar = document.getElementById("toolbar");
        const splitter = document.getElementById('splitter');
        if (sidebar.classList.contains("sidebar-close")) {
            splitter.classList.remove('hidden');
            sidebar.classList.remove("sidebar-close");
            sidebar.classList.add("sidebar-open");
            toolbar.classList.remove("toolbar-close");
            toolbar.classList.add("toolbar-open");
            let timer = setTimeout(() => {
                toolbar.classList.remove("toolbar-open");
                clearTimeout(timer);
            }, 201);
        } else {
            sidebar.classList.remove("sidebar-open");
            sidebar.classList.add("sidebar-close");
            toolbar.classList.remove("toolbar-open");
            toolbar.classList.add("toolbar-close");
            let timer = setTimeout(() => {
                splitter.classList.add('hidden');
                clearTimeout(timer);
            }, 208);
        }
    }
</script>

<div class="w-screen h-screen flex select-none">
    <div id="sidebar" class="h-full w-[{sidebarWidth}px]">
        <div id="sidetool" class="h-[39px] min-w-[110px] absolute flex items-center pl-[81px] z-50" style="width: {sidebarWidth}px">
            <SideLeading on:click={toggleSidebar} />
        </div>
        <div class="h-full w-full dark:text-white overflow-hidden flex flex-col" style="--wails-draggable:no-drag">
            <div class="flex-none h-[40px] border-b-[1px] border-solid border-gray-400 border-opacity-40 dark:border-opacity-20"></div>
            <div class="overflow-auto">
            </div>
        </div>
    </div>
    <div id="splitter" class="splitter splitter-horizontal cursor-col-resize w-[1px] bg-gray-400 bg-opacity-40 dark:bg-black dark:bg-opacity-40" style="--wails-draggable:no-drag"/>
    <div class="grow h-full">
        <div style="width: calc(100% - {sidebarWidth}px)" id="toolbar"
            class="h-[39px] pl-[10px] pr-[11px] flex items-center justify-between absolute right-0 top-0 z-50">
            <strong id="title" class="dark:text-white select-text" style="--wails-draggable:no-drag">Editor Header</strong>
            <SideLeading />
        </div>
        <div id="content" class="mt-[39px] h-full w-full border-t-[1px] border-solid border-gray-400 border-opacity-40 dark:border-opacity-20 dark:text-white select-text" style="--wails-draggable:no-drag">
            Content View
        </div>
    </div>
</div>

<style lang="scss">
    .sidebar-close {
        animation-name: sidebar-frams-close;
        animation-duration: 0.2s;
        animation-fill-mode: forwards;
        animation-timing-function: ease-in;
    }
    .sidebar-open {
        animation-name: sidebar-frams-open;
        animation-duration: 0.2s;
        animation-fill-mode: forwards;
        animation-timing-function: ease-in;
    }
    @keyframes sidebar-frams-open {
        from {
            width: 0px;
        }
    }
    @keyframes sidebar-frams-close {
        to {
            width: 0px;
        }
    }

    .toolbar-close {
        animation-name: toolbar-frams-close;
        animation-duration: 0.13s;
        animation-fill-mode: forwards;
        animation-timing-function: ease-in;
    }
    .toolbar-open {
        animation-name: toolbar-frams-open;
        animation-duration: 0.13s;
        animation-fill-mode: forwards;
        animation-timing-function: ease-in;
    }
    @keyframes toolbar-frams-close {
        to {
            width: calc(100% - 110px);
        }
    }
    @keyframes toolbar-frams-open {
        from {
            width: calc(100% - 110px);
        }
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
