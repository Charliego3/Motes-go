<script>
    import LeadingBtn from "./LeadingBtn.svelte";
    import {cubicInOut} from 'svelte/easing';
    import {tweened} from 'svelte/motion';
    import {onMount} from "svelte";
    import {FetchFileTree, Open} from '../wailsjs/go/main/App.js';
    import FileTree from "./FileTree.svelte";
    import TrailingBtn from "./TrailingBtn.svelte";
    import {EditorView, basicSetup} from "codemirror";
    import {markdown} from "@codemirror/lang-markdown";
    import {languages} from "@codemirror/language-data";

    export let sidebarWidth = 300;
    export let sidebarMinWidth = 200;
    export let sidebarMaxWidth = null;
    let dragging = false;
    const motionSidebarWidth = tweened(sidebarWidth, {
        duration: 200,
        easing: cubicInOut,
    });

    $: editorWidth = $motionSidebarWidth;
    $: if ($motionSidebarWidth <= 110) {
        editorWidth = 110;
    }

    let tree= [];
    let editor;
    let editorView;
    onMount(() => {
        FetchFileTree().then((files) => {
            tree = files;
        }).catch((err) => { 
            console.log(err);
        })

        // let state = EditorState.create({
        //     doc: editorContent,
        // });
        editorView = new EditorView({
            extensions: [basicSetup, markdown({codeLanguages: languages}), EditorView.lineWrapping, EditorView.theme({
                "&": {
                    // color: "white",
                    backgroundColor: "transparent",
                    outline: "none !important",
                    // width: "calc(100% - "+editorWidth+"px)",
                },
                ".cm-scroller": {overflow: "auto"},
                ".ͼ1.cm-editor.cm-focused": {
                    outline: "none"
                },
                ".cm-content": {
                    caretColor: "#0e9"
                },
                "&.cm-focused .cm-cursor": {
                    // borderLeftColor: "#0e9"
                    width: "5px"
                },
                "&.cm-focused .cm-selectionBackground, ::selection": {
                    backgroundColor: "dodgerblue"
                },
                ".cm-gutters": {
                    backgroundColor: "transparent",
                    color: "#ddd",
                    borderRight: "0px",
                },
                ".cm-lineNumbers": {
                    color: "gray",
                    userSelect: "none !important",
                },
                ".cm-activeLineGutter": {
                    // backgroundColor: "#006cc3",
                    backgroundColor: "transparent",
                    color: "#3283E4"
                },
                ".cm-activeLine": {
                    // color: "white",
                    backgroundColor: "#99999945",
                    borderRadius: "3px",
                },
            })],
            parent: editor,
        });

        document.addEventListener('mouseup', function () {
            dragging = false;
            document.removeEventListener('mousemove', mouseMoving);
        });
    });

    function mouseMoving(e) {
        let width = e.pageX;
        if (width <= sidebarMinWidth) {
            return;
        }
        if ((sidebarMaxWidth && width >= sidebarMaxWidth) || (width >= document.body.clientWidth / 2)) {
            return;
        }
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

    let selectedNode;
    let selectedItem;
    let editorHeader = '';
    function chooseFile(node, isDir) {
        if (!!selectedNode) {
            selectedNode.classList.remove('selected');
        }
        if (!!selectedItem) {
            selectedItem.classList.remove('selected-folder');
            selectedItem.classList.add('unselected-folder');
        }
        selectedNode = node;
        selectedNode.classList.add('selected');

        let path;
        if (isDir) {
            path = selectedNode.getElementsByTagName('path')[1];
        } else {
            path = selectedNode.getElementsByTagName('path')[0];
        }
        console.log(isDir, path);
        selectedItem = path;
        selectedItem.classList.remove('unselected-folder');
        selectedItem.classList.add('selected-folder');
    }

    let editorContent = '';
    function open(name, fpath) {
        editorHeader = name;
        console.log(name, fpath);
        Open(fpath).then(content => {
            // editorContent = content;
            let transaction = editorView.state.update({
                changes: {from: 0, to: editorView.state.doc.length, insert: content}
            });
            editorView.dispatch(transaction);
        }).catch(err => {
            console.log("open file error", err);
        })
    }
</script>

<div class="w-screen h-screen flex select-none">
    <div id="sidebar" class="h-full" style="width: {$motionSidebarWidth}px;">
        <div id="sidetool" class="h-[39px] min-w-[110px] absolute flex items-center pl-[81px] z-50"
             style="width: {$motionSidebarWidth}px">
            <LeadingBtn on:click={toggleSidebar}/>
        </div>
        <div class="h-full w-full dark:text-white flex flex-col" style="--wails-draggable:no-drag;">
            <div class="flex-none h-[39px] border-b-[1px] border-solid border-gray-400 border-opacity-40 dark:border-opacity-20"></div>
            <div class="overflow-y-scroll overflow-x-hidden whitespace-nowrap py-[10px] pl-[10px] scrollbar"
                 style="padding-right: {$motionSidebarWidth - 50 <= 0 ? 0 : 10}px;width: {$motionSidebarWidth}px;">
                <FileTree {open} children={tree} onClick={chooseFile} expanded/>
            </div>
        </div>
    </div>
    <div id="splitter" on:mousedown={() => {dragging = true;document.addEventListener('mousemove', mouseMoving)}}
         class:hidden={$motionSidebarWidth === 0}
         class="splitter splitter-horizontal cursor-col-resize w-[1px] bg-gray-400 bg-opacity-30 dark:bg-gray-500 dark:bg-opacity-40"
         style="--wails-draggable:no-drag"></div>
    <div class="grow h-full">
        <div style='width: calc(100% - {editorWidth+1}px);' id="toolbar"
             class="h-[39px] pl-[10px] pr-[11px] flex items-center justify-between absolute right-0 top-0 z-50">
            <strong id="title" class:select-none={dragging} class="dark:text-white select-text"
                    style="--wails-draggable:no-drag;">{editorHeader}</strong>
            <div class="flex-none"><TrailingBtn/></div>
        </div>
        <div id="content" class:select-none={dragging}
             class="flex-grow-0 h-full w-full flex flex-col dark:text-white select-text"
             style="--wails-draggable:no-drag">
            <div class="flex-none h-[39px] border-b-[1px] border-solid border-gray-400 border-opacity-40 dark:border-opacity-20"></div>
            <div class="w-full h-full bg-transparent overflow-auto">
                <div bind:this={editor}></div>
                <!-- <div><Markdown md={editorContent} plugins={[gfmPlugin]}/></div> -->
                <!-- <div>{@html editorContent}</div> -->
            </div>
        </div>
    </div>
</div>

<style lang="scss">
    :global(.scrollbar) {
        scrollbar-3dlight-color: transparent;
    }

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
            // background-color: aquamarine;
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
