<script setup lang="ts">
// problème -> le bouton play ne veut pas se mettre en état disabled
const props = defineProps(['animationRunning']);
const emit = defineEmits<{
    click: [];
}>();
</script>

<!-- Wraps any element in a zero-padding button -->
<template>
    <button class="w-min h-min" @click="emit('click')" :disabled="animationRunning">
        <slot></slot>
    </button>
</template>
