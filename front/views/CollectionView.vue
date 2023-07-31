<script setup lang="ts">
import TextualButton from '@/components/inputs/TextualButton.vue';
import GraphQL from '@/components/GraphQL.vue';
import { GetSongDocument, GetSongQuery } from '@/gql/queries/GetSong';
import type { ApolloQueryResult } from '@apollo/client/core/types';
import { useRouter } from 'vue-router';
import { useSongStore } from '@/store';
import { Song } from '@/gql/types';

const router = useRouter();
const songStore = useSongStore();

const handleClick = (song: Partial<Song>) => {
    songStore.setCurrentSong(song);

    router.push({
        path: `/collection/${song?.url}`,
    });
};
</script>

<template>
    <GraphQL :query="GetSongDocument">
        <template #default="{ data }: ApolloQueryResult<GetSongQuery>">
            <div class="container mx-auto px-4 lg:px-0 lg:max-w-4xl">
                <main v-if="data" title="collection section" class="mt-16">
                    <h1 class="text-center text-3xl font-semibold m-5">Collection</h1>
                    <div
                        v-for="(item, index) in data.songs"
                        :key="index"
                        class="mx-4 lg:mx-0 grid grid-cols-2 gap-4 items-center"
                    >
                        <div class="col-span-1">
                            <h2 class="text-center mt-2">{{ item.title }}</h2>
                            <img style="width: 300px" :src="item?.img_url" :alt="item.title" class="" />
                        </div>
                        <div class="col-span-1">
                            <TextualButton
                                aria-label="example button G"
                                hover-color="#FAFF00"
                                @click="handleClick(item)"
                            >
                                Lancer {{ item.title }}
                            </TextualButton>
                        </div>
                    </div>
                </main>
            </div>
        </template>
    </GraphQL>
</template>
