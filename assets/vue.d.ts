// Dummy declaration to satisfy the import requirements of WebStorm/GoLand

declare module '*.vue' {
    import type { ComponentOptions } from 'vue';

    declare const component: ComponentOptions;
    export default component;
}
