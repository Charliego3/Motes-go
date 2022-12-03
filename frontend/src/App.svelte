<script>
    import SideLeading from "./icons/SideLeading.svelte";
    import { Pane, Splitpanes } from "svelte-splitpanes";

    const defaultWidth = 30;
    let sidebarWidth = defaultWidth;
    let sidebarOpened = true;
    $: editorWidth = 100 - sidebarWidth;

    function resized(e) {
        sidebarWidth = e.detail[0].size;
    }

    function toggleSidebar(e) {
        let sidebar = document.getElementsByClassName("root__panel")[0];
        let isOpen = sidebar.classList.contains("sidebar-close");
        sidebar.nextElementSibling.nextElementSibling.removeAttribute("style");
        let leading = document.getElementById('sidebarLeading');
        if (isOpen) {
            // open
            sidebarOpened = true;
            toggleSidebarSplitter(sidebar, true);
            sidebar.classList.remove("sidebar-close");
            sidebar.classList.add("sidebar-open");
            leading.classList.remove("sidebar-close");
            leading.classList.add('sidebar-open');
            let timer = setTimeout(() => {
                sidebar.nextElementSibling.nextElementSibling.setAttribute(
                    "style",
                    "width: " + editorWidth + "%"
                );
                sidebar.classList.remove("sidebar-open");
                leading.classList.remove("sidebar-open");
                clearTimeout(timer);
            }, 200);
        } else {
            // close sidebar
            sidebarOpened = false;
            sidebar.classList.remove("sidebar-open");
            sidebar.classList.add("sidebar-close");
            leading.classList.remove("sidebar-open");
            leading.classList.add("sidebar-close");
            let timer = setTimeout(() => {
                toggleSidebarSplitter(sidebar, false);
                clearTimeout(timer);
            }, 200);
        }
    }

    function toggleSidebarSplitter(sidebar, opened) {
        let splitter = sidebar.nextElementSibling;
        if (opened) {
            splitter.classList.remove("hidden");
        } else {
            splitter.classList.add("hidden");
        }
    }
</script>

<main class="w-screen h-screen">
    <!-- Toolbar View Start -->
    <div
        id="toolbar"
        class="absolute top-0 h-[39px] 
    flex gap-2 items-center z-50 w-full border-b border-solid border-[#C7C7C8] dark:border-[#565557]
    "
    >
        <div id="sidebarLeading" style="width: {sidebarOpened ? sidebarWidth : 0}%;" class="pl-[81px] pr-[35px]">
            <SideLeading on:click={toggleSidebar} />
        </div>
        <div class="flex justify-between w-full pr-[10px]" style="{sidebarOpened ? 'width: ' + sidebarWidth + '%' : ''}"> 
            <span style="--wails-draggable:no-drag;">Editor Header</span>
            <SideLeading />
        </div>
    </div>
    <!-- Toolbar View Ended -->

    <!-- Split Content View Start -->
    <div class="w-screen h-screen" style="--wails-draggable:no-drag">
        <Splitpanes
            class="text-black dark:text-[#ffffff]"
            horizontal={false}
            theme="root-theme"
            on:resize={resized}
        >
            <!-- Sidebar View Start -->
            <Pane
                class={"w-[300px] root__panel resize-none"}
                size={sidebarWidth}
                minSize={20}
                maxSize={50}
            >
                <div class="pt-[39px] w-full h-full select-none">flex 2</div>
            </Pane>
            <!-- Sidebar View Ended -->

            <!-- Editor View Start -->
            <Pane
                class="min-w-[50%] pt-[39px] max-w-full select-none dark:bg-[#292A2F] bg-[#FFFFFF]"
                maxSize={100}
            >
                <span
                    >3

                    <br />
                    <em class="specs">I have a max height of 70%</em>
                    <code>
                        // // Login.swift // SpotifyMenu (iOS) // // Created by
                        Charlie on 2022/11/26. // // import SwiftUI // struct
                        Login: View // //Environment(\.openURL) private var
                        openURL //@EnvironmentObject var network: SpotifyNetwork
                        //var body: some View // ZStack // Image("Banner") //
                        .opacity(0.5) // .resizable() //
                        .aspectRatio(contentMode: .fill) // // .frame(width:
                        Contants.width, height: Contants.height) // .padding()
                        // // // struct Login_Previews: PreviewProvider //
                        static var previews: some View // Login() // //
                    </code>
                </span>
            </Pane>
            <!-- Editor View Ended -->
        </Splitpanes>
        <!-- Split Content View Ended -->
    </div>
</main>

<style global lang="scss">
    .sidebar-close {
        animation-name: sidebar-slide;
        animation-duration: 0.2s;
        animation-fill-mode: forwards;
        animation-timing-function: ease-in;
    }

    .sidebar-open {
        animation-name: sidebar-slide-open;
        animation-duration: 0.2s;
        animation-fill-mode: forwards;
        animation-timing-function: ease-in;
    }

    @keyframes sidebar-slide-open {
        from {
            width: 0px;
        }
    }

    @keyframes sidebar-slide {
        to {
            width: 0%;
        }
    }

    .splitpanes.root-theme {
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
                background-color: #d3d3d4;
            }
        }

        @media (prefers-color-scheme: dark) {
            .splitpanes__splitter {
                background-color: #1d1e1f;
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
