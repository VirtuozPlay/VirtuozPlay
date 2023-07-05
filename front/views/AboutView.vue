<script setup lang="ts">
import CardApp from '../components/CardApp.vue';
import VirtuozPlayVersion from '@/components/VirtuozPlayVersion.vue';
import SquareIcon from '@/components/icons/SquareIcon.vue';
import SimpleButton from '@/components/inputs/SimpleButton.vue';
import TextualButton from '@/components/inputs/TextualButton.vue';
</script>

<template>
    <main aria-label="about section" class="about items-center mt-16 w-80vw flex justify-center">
        <CardApp class="font-bold">
            <VirtuozPlayVersion />
            <h3>Crédit musique</h3>
            <p>Partie sur les crédits musique</p>
        </CardApp>
        <div class="flex flex-col gap-2">
            <p class="font-bold">Exemples de boutons</p>
            <SimpleButton aria-label="example button A">
                <SquareIcon :icon="['fas', 'bars']" hover-color="#FAFF00" />
            </SimpleButton>
            <SimpleButton aria-label="example button B">
                <SquareIcon :icon="['fas', 'arrow-right']" color="#FAFF00" hover-color="#FAFFCA" />
            </SimpleButton>
            <SimpleButton aria-label="example button C">
                <SquareIcon :icon="['fas', 'arrow-right']" color="#00FF75" hover-color="#FAFF00" />
            </SimpleButton>
            <SimpleButton aria-label="example button D">
                <SquareIcon :icon="['fas', 'arrow-right']" color="#FFF" hover-color="#DDD" />
            </SimpleButton>
            <SimpleButton aria-label="example button E">
                <SquareIcon :icon="['fas', 'user-secret']" color="#FAFF00" hover-color="#FAFFCA" />
            </SimpleButton>
            <SimpleButton aria-label="example button F">
                <SquareIcon :icon="['fas', 'play']" color="#00FF75" hover-color="#FAFF00" />
            </SimpleButton>
            <TextualButton aria-label="example button G" hover-color="#FAFF00">Exemple</TextualButton>
        </div>
    </main>
</template>
