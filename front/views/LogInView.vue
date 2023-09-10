<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useUserStore } from '@/store';
import router from '@/router';

const userStore = useUserStore();

const email = ref('');
const password = ref('');
const errors: string[] = reactive([]);

async function onLogin(event: Event) {
    event.preventDefault();

    errors.slice(0, errors.length);
    errors.push(...(await userStore.logIn(email.value, password.value)));
    if (errors.length === 0) {
        await router.push({ name: 'home' });
    }
}
</script>

<template>
    <h1 class="text-2xl">Se connecter à VirtuozPlay</h1>
    <form class="flex flex-col w-1/3 gap-2" @submit.prevent="onLogin">
        <label for="email">Email</label>
        <input id="email" v-model="email" type="text" placeholder="Email" />
        <label for="password">Mot de passe</label>
        <input id="password" v-model="password" type="password" placeholder="Mot de passe" />

        <ul v-if="errors.length > 0" class="text-red-700 font-bold">
            <li v-for="error in errors" :key="error">{{ error }}</li>
        </ul>

        <button type="submit">Se connecter</button>
        <RouterLink :to="{ name: 'signup' }">Pas encore de compte?</RouterLink>
    </form>
</template>
