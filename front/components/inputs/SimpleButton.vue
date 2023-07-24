<script setup lang="ts">
// problème -> le bouton play ne veut pas se mettre en état disabled
defineProps<{
    animationRunning?: boolean;
}>();
const emit = defineEmits<{
    click: [];
}>();
</script>

<!-- Wraps any element in a zero-padding button -->
<template>
    <button class="w-min h-min" :disabled="animationRunning" @click="emit('click')">
        <slot></slot>
    </button>
</template>
