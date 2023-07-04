<script setup lang="ts">
import { RouterLink } from 'vue-router';
import SquareIcon from '@/components/icons/SquareIcon.vue';
import SimpleButton from '@/components/inputs/SimpleButton.vue';
</script>

<template>
    <div
        class="flex flex-wrap items-center justify-between w-screen bg-white border-b-4 border-solid border-primary-text"
    >
        <div class="p-4">
            <img alt="VirtuozPlay Logo" src="@/assets/virtuozplay-logo.png" width="125" height="125" />
        </div>
        <nav class="m-8 flex items-center flex-wrap p-4 gap-8">
            <RouterLink to="/about" class="text-xl font-bold">A Propos</RouterLink>
            <RouterLink to="/profile" class="text-xl font-bold">Profil</RouterLink>
            <RouterLink class="noeffect" to="/">
                <SimpleButton>
                    <SquareIcon :icon="['fas', 'bars']"></SquareIcon>
                </SimpleButton>
            </RouterLink>
        </nav>
    </div>
</template>
