<script setup lang="ts">
import SquareIcon from '@/components/icons/SquareIcon.vue';

export interface Props {
    to: string;
    selected?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
    selected: false,
});

const emit = defineEmits<{
    select: [];
}>();
</script>

<template>
    <div @mouseover="emit('select')" @focusin="emit('select')" class="relative">
        <SquareIcon
            v-show="props.selected"
            :icon="['fas', 'arrow-right']"
            color="#FAFF00"
            class="absolute left-[-100px]"
        ></SquareIcon>
        <RouterLink :to="props.to" class="hover:bg-link-selection focus:bg-link-selection outline-none">
            <h2 class="font-extrabold text-7xl bg-inherit rounded-full transition-colors duration-200 ease-in">
                <slot name="default" />
            </h2>
            <h3 class="text-3xl">
                <slot name="subtitle" />
            </h3>
        </RouterLink>
    </div>
</template>
