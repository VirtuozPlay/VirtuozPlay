<script setup lang="ts">
import { defineComponent } from 'vue';
import CardApp from '../components/CardApp.vue';
import ButtonApp from '../components/ButtonApp.vue';

defineComponent({
    components: {
        CardApp,
        ButtonApp,
    },
});
</script>

<template>
    <div class="profil main items-center">
        <CardApp>
            <div class="flex">
                <div class="avatar">
                    <font-awesome-icon icon="fa-solid fa-user-secret" size="3x" />
                </div>
                <h1>toto</h1>
            </div>
            <div>
                <ButtonApp>Modifier mon profil</ButtonApp>
            </div>
        </CardApp>
    </div>
</template>

<style>
.avatar {
    color: #4c324d;
    background-color: var(--color);
    width: 105px;
    height: 105px;
    padding: 10px 8px 8px 15px;
    font-size: 25px;
    border-radius: 4px;
    font-weight: 700;
    border: 4px solid #4c324d;
    box-shadow: 8px 8px #4c324d;
    border-radius: 50px;
    transition: background-color 0.3s ease;
    margin-right: 20px;
}
</style>
