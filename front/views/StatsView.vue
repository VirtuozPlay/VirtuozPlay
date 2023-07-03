<script setup lang="ts">
import { defineComponent } from 'vue';
import CardApp from '../components/CardApp.vue';

defineComponent({
    components: {
        CardApp,
    },
});
</script>

<template>
    <div class="profil main items-center">
        <CardApp>
            <div>
                <h1>Stats</h1>
            </div>
        </CardApp>
    </div>
</template>
