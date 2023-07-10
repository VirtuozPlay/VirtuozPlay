import { describe, expect, it } from 'vitest';

import { mount } from '@vue/test-utils';
import TextualButton from '@/components/inputs/TextualButton.vue';

describe('TextualButton', () => {
    it('renders properly', () => {
        const wrapper = mount(TextualButton, { slots: { default: 'Click Me!' } });

        expect(wrapper.text()).toContain('Click Me!');
    });

    it('handles click event', () => {
        let clickCount = 0;
        const Button = {
            components: { TextualButton },
            methods: {
                onClick() {
                    ++clickCount;
                },
            },
            template: '<TextualButton @click="onClick">HERE</TextualButton>',
        };
        const wrapper = mount(Button);

        wrapper.findComponent(TextualButton).trigger('click');
        expect(clickCount).toBe(1);
    });
});
