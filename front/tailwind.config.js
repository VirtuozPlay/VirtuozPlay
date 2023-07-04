/* eslint-env node */
/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./**/*.{vue,html,ts,js}'],
    theme: {
        extend: {
            boxShadow: {
                normal: '8px 8px 0 0 #4c324d',
            },
            colors: {
                'primary-text': '#4C324d',
                'link-selection': 'hsla(66, 100%, 55%, 0.6)',
            },
        },
    },
    plugins: [],
};
