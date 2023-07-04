<script setup lang="ts">
import HomeLink from '@/components/HomeLink.vue';
</script>

<template>
    <main class="w-screen grow flex">
        <nav class="flex flex-col w-min h-1/2 m-auto gap-4 justify-evenly">
            <HomeLink to="/collection">
                Collection
                <template #subtitle>Choisir un morceau</template>
            </HomeLink>
            <HomeLink to="/profile">
                Profil
                <template #subtitle>Mes informations</template>
            </HomeLink>
            <HomeLink to="/stats">
                Stats
                <template #subtitle>Mes performances</template>
            </HomeLink>
            <HomeLink to="/checkup">
                Checkup
                <template #subtitle>Accorder ma guitare</template>
            </HomeLink>
        </nav>
    </main>
</template>
