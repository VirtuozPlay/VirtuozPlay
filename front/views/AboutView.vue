<script>
import Card from '../components/Card.vue';
import HelloWorld from '../components/HelloWorld.vue';
import Button from '../components/Button.vue';
import Icon from '../components/Icon.vue';
const handleClick = () => {
    console.log('Bouton cliqué !');
};

export default {
    components: {
        Card,
        Button,
        Icon,
    },
};
</script>

<template>
    <div class="about main items-center">
        <Card>
            <HelloWorld msg="Hi, VirtuozPlay" />
            <div>
                <h1>About</h1>
                <p>VirtuozPlay..</p>
                <h3>Crédit musique</h3>
                <p>Partie sur les crédits musique</p>
            </div>
        </Card>
        <div>
            <p>Exemples de bouton</p>
            <Icon icon="bars"></Icon>
            <Icon icon="arrow-right" color="#FAFF00"></Icon>
            <Icon icon="arrow-right" color="#00FF75"></Icon>
            <Icon icon="arrow-right" color="#FFF"></Icon>
            <Icon icon="user-secret" color="#FAFF00" />
            <Icon icon="play" color="#00FF75" />
            <Button @click="handleClick">Cliquez ici</Button>
        </div>
    </div>
</template>
