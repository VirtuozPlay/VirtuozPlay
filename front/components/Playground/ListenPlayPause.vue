<script setup lang="ts">
import SquareIcon from '@/components/icons/SquareIcon.vue';
import SimpleButton from '@/components/inputs/SimpleButton.vue';

const props = defineProps(['onListen', 'onPlay', 'onPause', 'onStop', 'animationRunning']);
</script>

<template>
    <main aria-label="about section" class="about items-center mt-16 w-80vw flex justify-center">
        <div class="flex flex-warp space-x-4">
            <SimpleButton :disabled="animationRunning" aria-label="" @click="onListen">
                <SquareIcon :icon="['fas', 'headphones']" color="#FAFF00" hover-color="#DDD" />
            </SimpleButton>
            <SimpleButton aria-label="" @click="onPlay">
                <SquareIcon :icon="['fas', 'play']" color="#00FF75" hover-color="#DDD" />
            </SimpleButton>
            <SimpleButton aria-label="" @click="onPause">
                <SquareIcon :icon="['fas', 'pause']" color="#C2C1FF" hover-color="#FAFF00" />
            </SimpleButton>
            <SimpleButton aria-label="" @click="onStop">
                <SquareIcon :icon="['fas', 'stop']" color="#FFCC15" hover-color="#FAFF00" />
            </SimpleButton>
        </div>
    </main>
</template>
