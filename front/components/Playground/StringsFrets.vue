<template>
    <div v-for="row in 6" :key="row" class="flex-container">
        <div v-for="column in 14" :key="column" class="flex-item">
            <!-- r {{ row }}- c{{ column }} -->
            <svg
                v-if="isPosition(row, column)"
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

<script setup lang="ts">
const props = defineProps(['isPosition']);
</script>

<style>
.flex-container {
    display: flex;
    flex-flow: row wrap;
    width: 100%;
    justify-content: space-around;
}

.flex-item {
    color: white;
    font-size: 0.5em;
    text-align: center;
    height: 52px;
    width: 100px;
    margin-top: -26px;
}

.anim {
    animation-duration: 1s;
    animation-name: tidou;
    transform-origin: 50px 26px;
    animation-iteration-count: 1;
    fill: #33ff8c;
}

@keyframes tidou {
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
