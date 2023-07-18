<script setup lang="ts">
import SimpleButton from '@/components/inputs/SimpleButton.vue';
import SquareIcon from '@/components/icons/SquareIcon.vue';
import { computed, ref } from 'vue';

const open = ref(false);

const backgroundStyle = computed(() => ({
    backgroundColor: open.value ? 'rgba(0, 0, 0, 0.4)' : 'transparent',
}));

const drawerStyle = computed(() => ({
    transform: `translateX(${open.value ? '0' : '240'}px)`,
    opacity: open.value ? '1' : '0',
}));
</script>

<template>
    <SimpleButton @click="open = true">
        <SquareIcon :icon="['fas', 'bars']" size="lg"></SquareIcon>
    </SimpleButton>
    <div
        v-show="open"
        id="drawer-background"
        aria-hidden="true"
        :style="backgroundStyle"
        class="inset-0 z-10 fixed"
        @click="open = false"
    ></div>
    <div
        aria-label="drawer"
        :style="drawerStyle"
        class="bg-white flex flex-col items-center flex-wrap p-4 gap-8 fixed right-0 top-0 bottom-0 z-20 transition ease-in border-4 border-r-0 border-solid border-primary-text duration-100 w-min-[240px] rounded-l-md"
    >
        <div class="w-full">
            <SimpleButton @click="open = false">
                <SquareIcon :icon="['fas', 'arrow-right']" size="lg" />
            </SimpleButton>
        </div>
        <slot name="default" />
    </div>
</template>
