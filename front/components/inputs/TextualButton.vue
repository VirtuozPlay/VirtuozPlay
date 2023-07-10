<script setup lang="ts">
import SimpleButton from '@/components/inputs/SimpleButton.vue';
import { ref } from 'vue';

export interface Props {
    color?: string;
    hoverColor?: string;
}

const props = withDefaults(defineProps<Props>(), {
    color: '#00ff75',
    hoverColor: undefined,
});

const bgColor = ref(props.color);
</script>

<template>
    <SimpleButton
        :style="{ backgroundColor: bgColor }"
        class="text-primary-text cursor-pointer button border-4 border-solid border-primary-text shadow-normal p-2 text-2xl font-bold rounded whitespace-nowrap transition-colors ease-out duration-100"
        @mouseover="bgColor = props.hoverColor || props.color"
        @mouseout="bgColor = props.color"
    >
        <slot></slot>
    </SimpleButton>
</template>
