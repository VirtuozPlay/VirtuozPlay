<script setup lang="ts">
import { defineComponent, ref } from 'vue';
import IconApp from '../components/IconApp.vue';

const isHovering1 = ref(false);
const isHovering2 = ref(false);
const isHovering3 = ref(false);
const isHovering4 = ref(false);

defineComponent({
    components: {
        IconApp,
    },

    setup() {
        return {
            isHovering1,
            isHovering2,
            isHovering3,
            isHovering4,
        };
    },
});
</script>

<template>
    <div class="about items-center flex-col mt-16 w-80vw flex justify-center">
        <div class="p-4">
            <div class="title flex">
                <IconApp v-if="isHovering4" icon="arrow-right icon" color="#FAFF00" class="inline-block mr-2"></IconApp>
                <h1 @mouseover="isHovering4 = true" @mouseout="isHovering4 = false">
                    <RouterLink to="/collection">Collection</RouterLink>
                </h1>
            </div>
            <h3>Choisir un morceaux</h3>
        </div>
        <div class="p-4">
            <div class="title flex">
                <IconApp v-if="isHovering1" icon="arrow-right icon" color="#FAFF00" class="inline-block mr-2"></IconApp>
                <h1 @mouseover="isHovering1 = true" @mouseout="isHovering1 = false">
                    <RouterLink to="/profil">Profil</RouterLink>
                </h1>
            </div>
            <h3>Mes informations</h3>
        </div>
        <div class="p-4">
            <div class="title flex">
                <IconApp v-if="isHovering2" icon="arrow-right icon" color="#FAFF00" class="inline-block mr-2"></IconApp>
                <h1 @mouseover="isHovering2 = true" @mouseout="isHovering2 = false">
                    <RouterLink to="/stats">Stats</RouterLink>
                </h1>
            </div>
            <h3>Mes performances</h3>
        </div>
        <div class="p-4">
            <div class="title flex">
                <IconApp v-if="isHovering3" icon="arrow-right icon" color="#FAFF00" class="inline-block mr-2"></IconApp>
                <h1 @mouseover="isHovering3 = true" @mouseout="isHovering3 = false">
                    <RouterLink to="/checkup">Checkup</RouterLink>
                </h1>
            </div>
            <h3>Accorder ma guitare</h3>
        </div>
    </div>
</template>

<style>
.title {
    position: relative;
    width: 400px;
    height: 50px;
}

.title > h1 {
    position: absolute;
    transform: translate(100px, -20px);
}

.title > .icon {
    position: absolute;
}

h3 {
    margin: 25px 0px 0px 102px;
}
</style>
