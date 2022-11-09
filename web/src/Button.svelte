<script lang="ts">

</script>


<button on:click>
    <slot></slot>
</button>

<style>
    button {
        @apply w-fit px-3.5 leading-loose bg-white/75 border rounded-md transition select-none;
    }

    button:hover {
        @apply bg-gray-100;
    }
</style>