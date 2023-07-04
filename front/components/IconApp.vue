<script setup lang="ts">
export interface Props {
    icon?: string;
    color?: string;
}

const props = withDefaults(defineProps<Props>(), {
    icon: '',
    color: '#C2C1FF',
});

const iconClass = `fa-${props.icon}`;

const handleClick = (): void => {
    console.log('Hello world');
};
</script>

<template>
    <div>
        <button :class="['icons', 'cursor-pointer']" :style="{ backgroundColor: color }" @click="handleClick">
            <font-awesome-icon :icon="iconClass" size="2x" />
        </button>
    </div>
</template>

<style lang="css" scoped>
.icons {
    color: #4c324d;
    background-color: var(--color);
    width: 72px;
    height: 72px;
    padding: 8px 8px;
    font-size: 25px;
    border-radius: 4px;
    font-weight: 700;
    border: 4px solid #4c324d;
    box-shadow: 8px 8px #4c324d;
    transition: background-color 0.3s ease;
}

.icons:hover {
    background-color: #faff00 !important;
    color: #4c324d;
}
</style>
