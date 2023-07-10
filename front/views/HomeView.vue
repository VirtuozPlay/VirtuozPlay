<script setup lang="ts">
import HomeLink from '@/components/HomeLink.vue';
import { ref } from 'vue';

/** Ensure only one link is active at a time, no matter what. */
const selection = ref(0);
</script>

<template>
    <main aria-label="home section" class="w-screen grow flex my-4">
        <nav
            aria-label="main menu"
            class="flex flex-col w-min h-1/2 m-auto gap-4 justify-evenly"
            @mouseout="selection = 0"
            @focusout="selection = 0"
        >
            <HomeLink to="/collection" :selected="selection === 1" @select="selection = 1">
                Collection
                <template #subtitle>Choisir un morceau</template>
            </HomeLink>
            <HomeLink to="/profile" :selected="selection === 2" @select="selection = 2">
                Profil
                <template #subtitle>Mes informations</template>
            </HomeLink>
            <HomeLink to="/stats" :selected="selection === 3" @select="selection = 3">
                Stats
                <template #subtitle>Mes performances</template>
            </HomeLink>
            <HomeLink to="/checkup" :selected="selection === 4" @select="selection = 4">
                Checkup
                <template #subtitle>Accorder ma guitare</template>
            </HomeLink>
        </nav>
    </main>
</template>
