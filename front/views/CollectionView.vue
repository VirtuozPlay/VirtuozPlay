<script setup lang="ts">
import TextualButton from '@/components/inputs/TextualButton.vue';
import GraphQL from '@/components/GraphQL.vue';
import { GetSongDocument, GetSongQuery } from '@/gql/queries/GetSong';
import type { QueryResult } from '@apollo/client';
import { useRouter } from 'vue-router';
import { useSongStore } from '@/store';
import { Song } from '@/gql/types';

const router = useRouter();
const songStore = useSongStore();

const handleClick = (song: Partial<Song>) => {
    songStore.setCurrentSong(song);

    router.push({
        path: `/collection/${song?.title}`,
    });
};
</script>

<template>
    <GraphQL :query="GetSongDocument">
        <template #default="{ data }: QueryResult<GetSongQuery>">
            <main v-if="data" title="collection section" class="mt-16 w-80vw mx-auto">
                <div class="w-full text-center"><h1>Collection</h1></div>
                <div v-for="(item, index) in data.songs" :key="index" class="mx-auto grid grid-cols-2 gap-4">
                    <div class="col-span-1">
                        <h2 class="text-center mt-2">{{ item?.title }}</h2>
                        <img src="https://placehold.co/600x400?text=Hello+World2" :alt="item?.title" class="w-full" />
                    </div>
                    <div class="col-span-1">
                        <TextualButton aria-label="example button G" hover-color="#FAFF00" @click="handleClick(item)">
                            Lancer {{ item?.title }}
                        </TextualButton>
                    </div>
                </div>
            </main>
        </template>
    </GraphQL>
</template>
