<script setup lang="ts">
import CardApp from '../components/CardApp.vue';
import AllTheNotes from '@/components/AllTheNotes.vue';
</script>

<template>
    <main aria-label="statistics section" class="items-center mt-16 w-80vw flex justify-center">
        <CardApp>
            <div>
                <h1>Stats</h1>
            </div>
            <AllTheNotes />
        </CardApp>
    </main>
</template>
