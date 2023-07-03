<script setup lang="ts">
import { defineComponent } from 'vue';
import CardApp from '../components/CardApp.vue';
import ButtonApp from '../components/ButtonApp.vue';
import IconApp from '../components/IconApp.vue';

defineComponent({
    components: {
        CardApp,
        ButtonApp,
        IconApp,
    },
});
</script>

<template>
    <div class="about main items-center">
        <CardApp>
            <HelloWorld msg="Hi, VirtuozPlay" />
            <div>
                <h1>About</h1>
                <p>VirtuozPlay..</p>
                <h3>Crédit musique</h3>
                <p>Partie sur les crédits musique</p>
            </div>
        </CardApp>
        <div>
            <p>Exemples de bouton</p>
            <IconApp icon="bars"></IconApp>
            <IconApp icon="arrow-right" color="#FAFF00"></IconApp>
            <IconApp icon="arrow-right" color="#00FF75"></IconApp>
            <IconApp icon="arrow-right" color="#FFF"></IconApp>
            <IconApp icon="user-secret" color="#FAFF00" />
            <IconApp icon="play" color="#00FF75" />
            <ButtonApp>Exemple</ButtonApp>
        </div>
    </div>
</template>
