import { describe, expect, it } from 'vitest';

import { mount } from '@vue/test-utils';
import SimpleButton from '@/components/inputs/SimpleButton.vue';

describe('SimpleButton', () => {
    it('renders properly', () => {
        const wrapper = mount(SimpleButton, { slots: { default: 'X' } });

        expect(wrapper.text()).toContain('X');
    });

    it('handles click event', () => {
        let clickCount = 0;
        const CloseButton = {
            components: { SimpleButton },
            methods: {
                onClick() {
                    ++clickCount;
                },
            },
            template: '<SimpleButton @click="onClick">X</SimpleButton>',
        };
        const wrapper = mount(CloseButton);

        wrapper.findComponent(SimpleButton).trigger('click');
        expect(clickCount).toBe(1);
    });
});
