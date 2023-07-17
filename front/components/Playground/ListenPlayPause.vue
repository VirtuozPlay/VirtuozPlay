<script setup lang="ts">
import SquareIcon from '@/components/icons/SquareIcon.vue';
import SimpleButton from '@/components/inputs/SimpleButton.vue';

const props = defineProps(['onListen', 'onPlay', 'onPause', 'onStop', 'animationRunning']);
</script>

<template>
    <main aria-label="about section" class="about items-center mt-16 w-80vw flex justify-center">
        <div class="flex flex-warp space-x-4">
            <SimpleButton @click="onListen" :disabled="animationRunning" aria-label="">
                <SquareIcon :icon="['fas', 'headphones']" color="#FAFF00" hover-color="#DDD" />
            </SimpleButton>
            <SimpleButton @click="onPlay" aria-label="">
                <SquareIcon :icon="['fas', 'play']" color="#00FF75" hover-color="#DDD" />
            </SimpleButton>
            <SimpleButton @click="onPause" aria-label="">
                <SquareIcon :icon="['fas', 'pause']" color="#C2C1FF" hover-color="#FAFF00" />
            </SimpleButton>
            <SimpleButton @click="onStop" aria-label="">
                <SquareIcon :icon="['fas', 'stop']" color="#FFCC15" hover-color="#FAFF00" />
            </SimpleButton>
        </div>
    </main>
</template>
