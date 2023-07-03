<script setup lang="ts">
defineProps<{
    msg: string;
}>();
</script>

<template>
    <div class="greetings">
        <h1 class="font-medium text-[2.6rem] relative top-[-10px] text-center lg:text-left">{{ msg }}</h1>
        <h3 class="font text-xl text-center lg:text-left"></h3>
    </div>
</template>
