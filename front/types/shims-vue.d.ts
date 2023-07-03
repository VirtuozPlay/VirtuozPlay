// Dummy declaration to satisfy the import requirements of WebStorm/GoLand

declare module '*.vue' {
    import type { DefineComponent } from 'vue';
    const component: DefineComponent;
    export default component;
}
