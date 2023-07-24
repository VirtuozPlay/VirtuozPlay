<script setup lang="ts">
import HeaderDrawer from '@/components/HeaderDrawer.vue';
import HeaderLink from '@/components/HeaderLink.vue';
</script>

<template>
    <header
        class="flex flex-wrap items-center justify-between w-screen bg-white border-b-4 border-solid border-primary-text"
    >
        <RouterLink :to="{ name: 'home' }">
            <img
                alt="VirtuozPlay Logo"
                src="@/assets/logo-128.png"
                class="h-[calc(48px+1rem)] md:h-[calc(48px+2rem)]"
            />
        </RouterLink>

        <nav aria-label="section navigation" class="flex items-center flex-wrap p-4 md:p-8 gap-8">
            <HeaderLink name="home" class="hidden md:inline">Accueil</HeaderLink>
            <HeaderLink name="profile" class="hidden md:inline">Profil</HeaderLink>
            <HeaderDrawer>
                <HeaderLink name="home">Accueil</HeaderLink>
                <HeaderLink name="profile">Profil</HeaderLink>
                <HeaderLink name="about">A Propos de VirtuozPlay</HeaderLink>
            </HeaderDrawer>
        </nav>
    </header>
</template>
