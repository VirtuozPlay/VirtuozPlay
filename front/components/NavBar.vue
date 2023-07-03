<script setup lang="ts">
import { defineComponent } from 'vue';
import { RouterLink } from 'vue-router';
import IconApp from './IconApp.vue';

defineComponent({
    name: 'NavBar',
    components: {
        RouterLink,
        IconApp,
    },
});
</script>

<template>
    <div class="flex flex-wrap items-center navbar justify-between">
        <div class="p-4">
            <img alt="VirtuozPlay Logo" src="@/assets/virtuozplay-logo.png" width="125" height="125" />
        </div>
        <nav class="m-8 flex items-center flex-wrap p-4">
            <RouterLink to="/about">About</RouterLink>
            <RouterLink class="noeffect" to="/">
                <div style="display: none"><font-awesome-icon icon="fa-solid fa-bars" size="3x" /></div>
                <IconApp icon="bars"></IconApp>
            </RouterLink>
        </nav>
    </div>
</template>

<style scoped>
nav a {
    font-size: 18px;
    padding: 0 1rem;
}

nav a:first-of-type {
    border: 0;
}

nav a.noeffect:hover {
    background-color: hsla(66, 100%, 55%, 0);
}
</style>
