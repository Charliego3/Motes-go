<script>
    import SideLeading from "../icons/SideLeading.svelte";

    export let sidebarWidth = 300;
    export let unit = "px";
    // let sidebarOpened = true;
    // let rightToolbarWidth = document.body.clientWidth - sidebarWidth;
    window.addEventListener("resize", function () {
        // rightToolbarWidth = document.body.clientWidth - sidebarWidth;
        // console.log(document.body.clientWidth, rightToolbarWidth);
        // document
        //     .getElementById("toolbar")
        //     .setAttribute(
        //         "style",
        //         "width: " + (document.body.clientWidth - sidebarWidth) + "px"
        //     );
    });
    function toggleSidebar() {
        let sidebar = document.getElementById("sidebar");
        let toolbar = document.getElementById("toolbar");
        let splitter = document.getElementById('splitter');
        if (sidebar.classList.contains("sidebar-close")) {
            splitter.removeAttribute('style');
            sidebar.classList.remove("sidebar-close");
            sidebar.classList.add("sidebar-open");
            toolbar.classList.remove("toolbar-close");
            toolbar.classList.add("toolbar-open");
            // sidebarOpened = true;
        } else {
            // sidebarOpened = false;
            sidebar.classList.remove("sidebar-open");
            sidebar.classList.add("sidebar-close");
            toolbar.classList.remove("toolbar-open");
            toolbar.classList.add("toolbar-close");
            let timer = setTimeout(() => {
                splitter.setAttribute('style', 'display: none;');
                clearTimeout(timer);
            }, 208);
        }
    }
</script>

<div class="w-screen h-screen flex select-none">
    <div id="sidebar" class="h-full w-[{sidebarWidth}{unit}]">
        <div
            class="h-[39px] min-w-[110px] absolute flex items-center pl-[81px] z-50"
        >
            <SideLeading on:click={toggleSidebar} />
        </div>
        <div class="pt-[39px] h-full dark:text-white bg-orange-500 overflow-hidden">
            <p class="bg-sky-700">Sidebar</p>
            <SideLeading on:click={toggleSidebar} />
        </div>
    </div>
    <div id="splitter"
        class="splitter splitter-horizontal cursor-col-resize w-[1px] bg-black dark:bg-[#1D1D1F]"
    />
    <div class="grow h-full">
        <div
            style="width: calc(100% - {sidebarWidth}{unit})"
            id="toolbar"
            class="h-[39px] pl-[10px] pr-[11px] w-full flex items-center justify-between absolute right-0 top-0"
        >
            <strong class="dark:text-white select-text" style="--wails-draggable:no-drag">Editor Header</strong>
            <SideLeading />
        </div>
        <div
            class="mt-[39px] border-t-[1px] border-solid border-black dark:text-white"
        >
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
            background-color: rgba(255, 0, 0, 0.3);
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

        &.splitter-vertical:before {
            top: -5px;
            bottom: -5px;
            width: 100%;
            cursor: col-resize;
        }
    }
</style>
