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
    <div class="items-center mt-16 w-80vw flex justify-center">
        <CardApp>
            <div>
                <h1>Stats</h1>
            </div>
        </CardApp>
    </div>
</template>
