<script setup lang="ts">
defineProps<{
    msg: string;
}>();
</script>

<template>
    <div class="greetings">
        <h1 class="green font-medium text-[2.6rem] relative top-[-10px] text-center lg:text-left">{{ msg }}</h1>
        <h3 class="font text-xl text-center lg:text-left">
            You’ve successfully created a project with
            <a href="https://vitejs.dev/" target="_blank" rel="noopener">Vite</a> +
            <a href="https://vuejs.org/" target="_blank" rel="noopener">Vue 3</a>. What's next?
        </h3>
    </div>
</template>
