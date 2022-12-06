<script>
    import SideLeading from "../LeadingBtn.svelte";
    import { Pane, Splitpanes } from "svelte-splitpanes";
    import Toolbar from "./Toolbar.svelte";

    export let sidebarWidth = 20;
    let previousSidevarWith = sidebarWidth;
    const sidebarMinWidthClass = "min-w-[217px]";

    function resized(e) {
        previousSidevarWith = sidebarWidth;
        sidebarWidth = e.detail[0].size;
    }

    async function toggleSidebar(e) {
        let sidebar = document.getElementsByClassName("root__panel")[0];
        if (sidebarWidth === 0) {
            sidebarWidth = previousSidevarWith;
            let timer = setTimeout(() => {
                sidebar.classList.add(sidebarMinWidthClass);
                clearTimeout(timer);
            }, 60);
        } else {
            sidebar.classList.remove(sidebarMinWidthClass);
            // sidebar.classList.add('w-[0px]');
            sidebar.removeAttribute('style');
            console.log("classList", sidebar.classList);
            sidebarWidth = 0;
            // let splitter = document.getElementsByClassName('root-theme')[0].getElementsByClassName('splitpanes__splitter')[0];
            // splitter.setAttribute('style', 'background-color: transparent;');
        }
    }
</script>

<Toolbar on:toggleSidebar={toggleSidebar} />
<div class="w-screen h-screen" style="--wails-draggable:no-drag">
    <Splitpanes
        class="text-black dark:text-[#ffffff]"
        horizontal={false}
        theme="root-theme"
        on:resized={resized}
    >
        <Pane
            class="min-w-[217px] hover:w-[0px] root__panel sidebarani"
        >

        <!-- bind:size={sidebarWidth} -->
            <div
                class="border-b border-solid flex items-center justify-end gap-[20px] border-[#C7C7C8] dark:border-[#565557] h-[39px] w-full pr-3"
                style="--wails-draggable:drag"
            >
                <SideLeading />
                <SideLeading />
            </div>
            <div class="w-full h-full select-none">flex 2</div>
        </Pane>
        <Pane class="min-w-[200px] pt-[38px] select-none">
            <span
                >3

                <br />
                <em class="specs">I have a max height of 70%</em>
            </span>
        </Pane>
    </Splitpanes>
</div>

<style global lang="scss">
    .sidebarani {
        transition: width 0.5s ease-out;
        -webkit-transition: width 0.5s ease-out;
    }

    .splitpanes.root-theme {
        .splitpanes__pane {
            background-color: transparent !important;
        }
        .splitpanes__splitter {
            position: relative;
            &:before {
                content: "";
                position: absolute;
                left: 0;
                top: 0;
                opacity: 0;
                z-index: 1;
                cursor: col-resize;
            }
            &.splitpanes__splitter__active {
                z-index: 2;
            }
        }
        .splitpanes__splitter:hover {
            cursor: col-resize;
        }

        @media (prefers-color-scheme: light) {
            .splitpanes__splitter {
                background-color: #c7c7c8;
            }
        }

        @media (prefers-color-scheme: dark) {
            .splitpanes__splitter {
                background-color: #565557;
            }
        }
    }

    .root-theme {
        &.splitpanes--vertical > .splitpanes__splitter:before {
            left: -5px;
            right: -5px;
            height: 100%;
            cursor: col-resize;
        }
    }
</style>
