<script setup lang="ts">
defineProps<{
    isPosition(string: number, fret: number): boolean;
}>();
</script>
<template>
    <div v-for="string in 6" :key="string" class="flex flex-row flex-wrap w-full justify-around">
        <div v-for="fret in 14" :key="fret" class="text-white text-sm text-center place">
            <svg
                v-if="isPosition(string, fret)"
                id="la_note"
                xmlns="http://www.w3.org/2000/svg"
                width="100"
                height="52"
                viewBox="0 0 100 52"
                role="img"
            >
                <circle stroke="#d0e0e3" stroke-miterlimit="10" class="anim" cx="50" cy="26" r="16" />
            </svg>
        </div>
    </div>
</template>

<style>
.place {
    height: 52px;
    width: 100px;
    margin-top: -26px;
}

.anim {
    animation-duration: 1s;
    animation-name: changecolor;
    transform-origin: 50px 26px;
    animation-iteration-count: 1;
    fill: #33ff8c;
}

@keyframes changecolor {
    0% {
        transform-origin: 40px 62px;
        fill: #f4a11a;
        transform: scale(0.8);
    }
    50% {
        transform-origin: 50px 26px;
        transform: scale(1);
    }
    70% {
        fill: #f4a11a;
    }
    80% {
        fill: #33ff8c;
        transform: scale(1);
    }
    100% {
        fill: #33ff8c;
        transform: scale(1);
    }
}
</style>
