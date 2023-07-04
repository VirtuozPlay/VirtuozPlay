<script setup lang="ts">
import NavBar from '@/components/NavBar.vue';
import { RouterView } from 'vue-router';
</script>

<template>
    <div id="app">
        <header>
            <NavBar />
        </header>
        <RouterView />
    </div>
</template>
