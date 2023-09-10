<script setup lang="ts">
import CardApp from '../components/CardApp.vue';
import ProfileIcon from '@/components/icons/ProfileIcon.vue';
import TextualButton from '@/components/inputs/TextualButton.vue';
import { useUserStore } from '@/store';
import router from '@/router';

const userStore = useUserStore();

async function logOut() {
    await userStore.logOut();
    await router.push({ name: 'home' });
}
</script>

<template>
    <main aria-label="profile section" class="items-center mt-16 w-80vw flex justify-center">
        <CardApp v-if="userStore.user === null">
            <h2 class="font-bold">Connectez-vous pour consuler votre profil</h2>
            <div class="flex flex-row gap-4">
                <TextualButton
                    hover-color="#FAFF00"
                    @click="() => router.push({ name: 'login', params: { redirectTo: 'home' } })"
                    >Se connecter</TextualButton
                >
                <TextualButton hover-color="#FAFF00" @click="() => router.push({ name: 'signup' })"
                    >Créer un compte</TextualButton
                >
            </div>
        </CardApp>
        <CardApp v-if="userStore.user !== null">
            <div class="flex items-center gap-2">
                <ProfileIcon />
                <h2 class="font-bold">{{ userStore.user.username }}</h2>
            </div>
            <div class="flex flex-row gap-4">
                <TextualButton hover-color="#FAFF00">Modifier mon profil</TextualButton>
                <TextualButton hover-color="#FAFF00" @click="logOut">Déconnexion</TextualButton>
            </div>
        </CardApp>
    </main>
</template>
