<script setup lang="ts">
import NavBar from '../assets/components/NavBar.vue';
import { RouterView } from 'vue-router';
</script>

<template>
    <div id="app">
        <header>
            <NavBar></NavBar>
        </header>
    </div>
    <RouterView />
</template>
