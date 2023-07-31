<script setup lang="ts">
defineProps<{
    name: string;
    value: unknown;
}>();
</script>

<template>
    <div class='border-4 border-solid border-primary-text w-max rounded-md shadow-normal p-2 inline-block'>
        <div class='text-md md:text-xl text-gray-600'>{{ name }}</div>
        <div class='text-2xl md:text-3xl text-red-900'>{{ value }}</div>
    </div>
</template>