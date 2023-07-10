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

    it('with base color only', () => {
        const Button = {
            components: { TextualButton },
            template: '<TextualButton color="#ff0000">HERE</TextualButton>',
        };
        const wrapper = mount(Button);

        const element = wrapper.findComponent(TextualButton).element as HTMLElement;
        expect(element.style.backgroundColor).toBe('rgb(255, 0, 0)');
        wrapper.findComponent(TextualButton).trigger('mouseover');
        wrapper.findComponent(TextualButton).trigger('mouseout');
    });

    it('with hover color', () => {
        const Button = {
            components: { TextualButton },
            template: '<TextualButton color="#ff0000" hover-color="#00ff00">HERE</TextualButton>',
        };
        const wrapper = mount(Button);

        const element = wrapper.findComponent(TextualButton).element as HTMLElement;
        expect(element.style.backgroundColor).toBe('rgb(255, 0, 0)');
        wrapper.findComponent(TextualButton).trigger('mouseover');
        wrapper.findComponent(TextualButton).trigger('mouseout');
    });
});
