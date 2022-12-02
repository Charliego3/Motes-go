<script>
    import SideLeading from "../icons/SideLeading.svelte";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    function toggleSidebar() {
        dispatch("toggleSidebar", {
            width: "Hello!",
        });
        // console.log("toggleSidebar")
    }
</script>

<div
    id="toolbar"
    class="absolute top-0 pl-[81px] h-[38px] content-center 
    flex gap-2 items-center z-50"
>
    <SideLeading on:click={toggleSidebar} />
</div>
