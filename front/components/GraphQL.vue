<script setup lang="ts">
// noinspection JSDeprecatedSymbols
import { ApolloQuery } from '@vue/apollo-components';
import type { DocumentNode } from 'graphql/language';
import type { ApolloQueryResult } from '@apollo/client/core/types';

withDefaults(
    defineProps<{
        query: DocumentNode;
        variables?: Record<string, unknown>;
        tag?: Element['tagName'];
    }>(),
    {
        variables: undefined,
        tag: 'div',
    }
);

// eslint-disable-next-line @typescript-eslint/no-explicit-any
defineSlots<ApolloQueryResult<any> & { default: any }>();
</script>

<!--
    This component is a wrapper around ApolloQuery that allows you to use the
    ApolloQuery component with TypeScript.
-->
<template>
    <ApolloQuery :query="() => $props.query" :variables="$props.variables" :tag='$props.tag'>
        <template #default="{ result }: { result: ApolloQueryResult<any> }">
            <slot v-bind="result"></slot>
        </template>
    </ApolloQuery>
</template>