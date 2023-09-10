<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useUserStore } from '@/store';
import router from '@/router';

const userStore = useUserStore();

const username = ref('');
const email = ref('');
const password = ref('');
const passwordConfirm = ref('');
const errors: string[] = reactive([]);

async function onSignup(event: Event) {
    event.preventDefault();

    if (password.value !== passwordConfirm.value) {
        errors.splice(0, errors.length);
        errors.push('Les mots de passe ne correspondent pas');
        return;
    }

    errors.slice(0, errors.length);
    errors.push(...(await userStore.signUp(username.value, email.value, password.value)));

    if (errors.length === 0) {
        await router.push({ name: 'home' });
    }
}
</script>

<template>
    <h1 class="text-2xl">Créer un compte VirtuozPlay</h1>
    <form class="flex flex-col w-1/3 gap-2" @submit.prevent="onSignup">
        <label for="username">Nom d'utilisateur</label>
        <input id="username" v-model="username" type="text" placeholder="Nom d'utilisateur" required />
        <label for="email">Email</label>
        <input id="email" v-model="email" type="email" placeholder="Email" required />
        <label for="password">Mot de passe</label>
        <input id="password" v-model="password" type="password" placeholder="Mot de passe" required />
        <label for="password">Confirmer mot de passe</label>
        <input id="passwordConfirm" v-model="passwordConfirm" type="password" placeholder="Confirmer" required />

        <ul v-if="errors.length > 0" class="text-red-700 font-bold">
            <li v-for="error in errors" :key="error">{{ error }}</li>
        </ul>

        <button type="submit">Créer un compte</button>
        <RouterLink :to="{ name: 'login' }">Déja inscrit?</RouterLink>
    </form>
</template>
