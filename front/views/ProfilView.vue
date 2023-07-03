<script>
import Card from '../components/Card.vue';
import Button from '../components/Button.vue';
import Icon from '../components/Icon.vue';
const handleClick = () => {
    console.log('Bouton cliqué !');
};

export default {
    name: 'Profil',
    components: {
        Card,
        Button,
        Icon,
    },
};
</script>

<template>
    <div class="profil main items-center">
        <Card>
            <div class="flex">
                <div class="avatar">
                    <font-awesome-icon icon="fa-solid fa-user-secret" size="3x" />
                </div>
                <h1>toto</h1>
            </div>
            <div>
                <Button @click="handleClick">Modifier mon profil</Button>
            </div>
        </Card>
    </div>
</template>

<style>
.avatar {
    color: #4c324d;
    background-color: var(--color);
    width: 105px;
    height: 105px;
    padding: 10px 8px 8px 15px;
    font-size: 25px;
    border-radius: 4px;
    font-weight: 700;
    border: 4px solid #4c324d;
    box-shadow: 8px 8px #4c324d;
    border-radius: 50px;
    transition: background-color 0.3s ease;
    margin-right: 20px;
}
</style>
