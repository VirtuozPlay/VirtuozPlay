declare module '@vue/apollo-components' {
    import type { DefineComponent, Plugin } from 'vue';

    /** @deprecated Please use the {@link @/GraphQL} component instead. */
    export const ApolloQuery: DefineComponent;

    export const VueApolloComponents: Plugin<unknown[]>;
    export default VueApolloComponents;
}
