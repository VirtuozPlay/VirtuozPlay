<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
import HelloWorld from './components/HelloWorld.vue';
</script>

<template>
    <header class="leading-normal max-h-screen lg:flex lg:place-items-center lg:pr-[calc(var(--section-gap)/2)]">
        <img
            alt="Vue logo"
            class="block mr-auto mb-8 lg:mr-8 lg:mb-0"
            src="@/assets/logo.svg"
            width="125"
            height="125"
        />

        <div class="flex place-items-start flex-wrap">
            <HelloWorld msg="You did it!" />

            <nav class="w-full font-size-[12px] text-center mt-8 lg:text-left lg:-ml-4 lg:text-base lg:py-4 lg:mt-4">
                <RouterLink to="/">Home</RouterLink>
                <RouterLink to="/about">About</RouterLink>
            </nav>
        </div>
    </header>

    <RouterView />
</template>

<style scoped>
/*noinspection CssUnusedSymbol*/
nav a.router-link-exact-active {
    color: var(--color-text);
}

/*noinspection CssUnusedSymbol*/
nav a.router-link-exact-active:hover {
    background-color: transparent;
}

nav a {
    display: inline-block;
    padding: 0 1rem;
    border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
    border: 0;
}
</style>
