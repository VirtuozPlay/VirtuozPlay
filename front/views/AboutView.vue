<script setup lang="ts">
import CardApp from '../components/CardApp.vue';
import VirtuozPlayVersion from '@/components/VirtuozPlayVersion.vue';
import SquareIcon from '@/components/icons/SquareIcon.vue';
import SimpleButton from '@/components/inputs/SimpleButton.vue';
import TextualButton from '@/components/inputs/TextualButton.vue';
</script>

<template>
    <div class="about items-center mt-16 w-80vw flex justify-center">
        <CardApp class="font-bold">
            <VirtuozPlayVersion />
            <h3>Crédit musique</h3>
            <p>Partie sur les crédits musique</p>
        </CardApp>
        <div class="flex flex-col gap-2">
            <p class="font-bold">Exemples de boutons</p>
            <SimpleButton>
                <SquareIcon :icon="['fas', 'bars']" hover-color="#FAFF00" />
            </SimpleButton>
            <SimpleButton>
                <SquareIcon :icon="['fas', 'arrow-right']" color="#FAFF00" hover-color="#FAFFCA" />
            </SimpleButton>
            <SimpleButton>
                <SquareIcon :icon="['fas', 'arrow-right']" color="#00FF75" hover-color="#FAFF00" />
            </SimpleButton>
            <SimpleButton>
                <SquareIcon :icon="['fas', 'arrow-right']" color="#FFF" hover-color="#DDD" />
            </SimpleButton>
            <SimpleButton>
                <SquareIcon :icon="['fas', 'user-secret']" color="#FAFF00" hover-color="#FAFFCA" />
            </SimpleButton>
            <SimpleButton>
                <SquareIcon :icon="['fas', 'play']" color="#00FF75" hover-color="#FAFF00" />
            </SimpleButton>
            <TextualButton hover-color="#FAFF00">Exemple</TextualButton>
        </div>
    </div>
</template>
