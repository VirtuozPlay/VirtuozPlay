<script setup lang="ts">
import { ref } from 'vue';
import SquareIcon from '@/components/icons/SquareIcon.vue';

export interface Props {
    to: string;
}

const props = defineProps<Props>();

const isHovering = ref(false);
</script>

<template>
    <div class="relative">
        <SquareIcon
            v-show="isHovering"
            :icon="['fas', 'arrow-right']"
            color="#FAFF00"
            class="absolute left-[-100px]"
        ></SquareIcon>
        <h1 class="text-7xl" @mouseover="isHovering = true" @mouseout="isHovering = false">
            <RouterLink
                :to="props.to"
                class="font-extrabold hover:bg-link-selection rounded-full transition-colors duration-200 ease-in"
            >
                <slot name="default" />
            </RouterLink>
        </h1>
        <h3 class="text-3xl">
            <slot name="subtitle" />
        </h3>
    </div>
</template>
