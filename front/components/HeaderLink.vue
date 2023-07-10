<script setup lang="ts">
import { RouterLink } from 'vue-router';

export interface Props {
    name: string;
}

defineProps<Props>();
</script>

<template>
    <RouterLink :to="{ name }" :class="{ underline: $route.name === name }" class="text-lg md:text-xl font-bold">
        <slot></slot>
    </RouterLink>
</template>
