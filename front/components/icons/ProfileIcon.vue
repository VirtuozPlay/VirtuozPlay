<script setup lang="ts">
import { FontAwesomeIcon, FontAwesomeIconProps } from '@fortawesome/vue-fontawesome';

export interface Props {
    size?: FontAwesomeIconProps['size'];
}

const props = withDefaults(defineProps<Props>(), {
    size: '4x',
});
</script>

<template>
    <FontAwesomeIcon
        :icon="['fas', 'user-secret']"
        :size="props.size"
        class="text-primary-text border-4 border-solid border-primary-text shadow-normal rounded-full w-[1em] h-[1em] p-[0.25em] m-4"
    />
</template>
