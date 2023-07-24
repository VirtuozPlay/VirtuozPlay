<script setup lang="ts">
// noinspection JSDeprecatedSymbols
import { ApolloQuery } from '@vue/apollo-components';
import type { DocumentNode } from 'graphql/language';
import type { QueryResult } from '@apollo/client';

defineProps<{
    query: DocumentNode;
    variables?: Record<string, unknown>;
}>();

// eslint-disable-next-line @typescript-eslint/no-explicit-any
defineSlots<QueryResult & { default: any }>();
</script>

<!--
    This component is a wrapper around ApolloQuery that allows you to use the
    ApolloQuery component with TypeScript.
-->
<template>
    <ApolloQuery :query="() => $props.query" :variables="$props.variables">
        <template #default="{ result }: { result: QueryResult }">
            <slot v-bind="result"></slot>
        </template>
    </ApolloQuery>
</template>
