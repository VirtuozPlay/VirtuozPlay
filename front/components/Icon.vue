<template>
    <div>
        <button :class="['icons', 'cursor-pointer']" :style="{ backgroundColor: color }" @click="handleClick">
            <font-awesome-icon :icon="iconClass" size="2x" />
        </button>
    </div>
</template>

<script>
export default {
    name: 'Icon',
    props: {
        icon: {
            type: String,
            required: true,
        },
        color: {
            type: String,
            default: '#C2C1FF',
        },
    },
    computed: {
        iconClass() {
            return `fa-${this.icon}`;
        },
    },
    methods: {
        handleClick() {
            console.log('hello world');
        },
    },
};
</script>

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
    border-radius: 4px;
    transition: background-color 0.3s ease;
}

.icons:hover {
    background-color: #faff00 !important;
    color: #4c324d;
}
</style>
