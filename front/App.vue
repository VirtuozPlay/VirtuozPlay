<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue';
import NavBar from '../assets/components/NavBar.vue';
</script>

<template>
    <div id="app">
        <header>
            <NavBar></NavBar>
        </header>
    </div>
    <RouterView />
</template>
