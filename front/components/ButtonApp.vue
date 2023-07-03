<template>
    <div>
        <button type="button" class="cursor-pointer button" @click="handleClick">
            <slot></slot>
        </button>
    </div>
</template>

<script setup lang="ts">
import { defineComponent } from 'vue';
defineComponent({
    name: 'ButtonApp',
});

const handleClick = (): void => {
    console.log('Hello world');
};
</script>

<style scoped>
.button {
    color: #4c324d;
    background-color: #00ff75;
    padding: 12px 24px;
    font-size: 25px;
    border-radius: 4px;
    font-weight: 700;
    border: 4px solid #4c324d;
    box-shadow: 8px 8px #4c324d;
    border-radius: 4px;
    transition: background-color 0.3s ease;
}

.button:hover {
    background-color: #faff00;
    color: #000;
}
</style>
