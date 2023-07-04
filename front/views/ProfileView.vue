<script setup lang="ts">
import CardApp from '../components/CardApp.vue';
import ProfileIcon from '@/components/icons/ProfileIcon.vue';
import TextualButton from '@/components/inputs/TextualButton.vue';
</script>

<template>
    <div class="items-center mt-16 w-80vw flex justify-center">
        <CardApp>
            <div class="flex items-center gap-2">
                <ProfileIcon />
                <h2 class="font-bold">Toto le gratteur</h2>
            </div>
            <TextualButton hover-color="#FAFF00">Modifier mon profil</TextualButton>
        </CardApp>
    </div>
</template>
