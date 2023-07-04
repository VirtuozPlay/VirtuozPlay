<script setup lang="ts">
const emit = defineEmits<{
    click: [];
}>();
</script>

<!-- Wraps any element in a zero-padding button -->
<template>
    <button class="w-min h-min" @click="emit('click')">
        <slot />
    </button>
</template>
