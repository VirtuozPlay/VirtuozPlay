<script setup lang="ts">
import TextualButton from '@/components/inputs/TextualButton.vue';
import { RouterLink } from 'vue-router';

interface Item {
    imageUrl: string;
    title: string;
    song: string;
}

defineProps({
    items: {
        type: Array as () => Item[],
        default: () => [
            {
                imageUrl: 'https://placehold.co/600x400?text=Hello+World1',
                title: 'Can Can',
                song: 'cancan',
            },
            {
                imageUrl: 'https://placehold.co/600x400?text=Hello+World3',
                title: 'Sting',
                song: 'sting',
            },
            {
                imageUrl: 'https://placehold.co/600x400?text=Hello+World2',
                title: 'Alabama',
                song: 'alabama',
            },
        ],
    },
});
</script>

<template>
    <main title="collection section" class="mt-16 w-80vw mx-auto">
        <div class="w-full text-center"><h1>Collection</h1></div>
        <div v-for="(item, index) in items" :key="index" class="mx-auto grid grid-cols-2 gap-4">
            <div class="col-span-1">
                <h2 class="text-center mt-2">{{ item.title }}</h2>
                <img :src="item.imageUrl" :alt="item.title" class="w-full" />
            </div>
            <div class="col-span-1">
                <RouterLink :to="`/${item.song}`">
                    <TextualButton aria-label="example button G" hover-color="#FAFF00">
                        Lancer {{ item.title }}
                    </TextualButton>
                </RouterLink>
            </div>
        </div>
    </main>
</template>
