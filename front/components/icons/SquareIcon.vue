<script setup lang="ts">
import { FontAwesomeIcon, FontAwesomeIconProps } from '@fortawesome/vue-fontawesome';
import { ref } from 'vue';

export interface Props {
    icon: FontAwesomeIconProps['icon'];
    size?: FontAwesomeIconProps['size'];
    color?: string;
    /** Color to use when hovering the icon, optional. */
    hoverColor?: string;
}

const props = withDefaults(defineProps<Props>(), {
    size: '2xl',
    color: '#C2C1FF',
    hoverColor: undefined,
});

const bgColor = ref(props.color);
</script>

<template>
    <FontAwesomeIcon
        :icon="props.icon"
        :size="props.size"
        :style="{ backgroundColor: bgColor }"
        class="icon text-primary-text border-4 border-solid border-primary-text w-[1em] h-[1em] p-[0.5em] rounded-md font-bold shadow-normal transition-colors ease-out duration-100"
        @mouseover="bgColor = props.hoverColor || props.color"
        @mouseout="bgColor = props.color"
    />
</template>
