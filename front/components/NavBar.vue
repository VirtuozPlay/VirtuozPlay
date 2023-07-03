<script>
import { RouterLink, RouterView } from 'vue-router';
import Icon from './Icon.vue';
export default {
    name: 'NavBar',
    components: {
        Icon,
    },
};
</script>
<template>
    <div class="flex flex-wrap items-center navbar justify-between">
        <div class="p-4">
            <img alt="VirtuozPlay Logo" src="@/assets/virtuozplay-logo.png" width="125" height="125" />
        </div>
        <nav class="m-8 flex items-center flex-wrap p-4">
            <RouterLink to="/about">About</RouterLink>
            <RouterLink class="noeffect" to="/">
                <font-awesome-icon icon="fa-solid fa-bars" size="3x" />
                <div style="display: none"><Icon icon="bars"></Icon></div>
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
