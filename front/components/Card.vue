<script></script>

<template>
    <div class="card">
        <div class="card-content">
            <slot></slot>
        </div>
    </div>
</template>

<style lang="css" scoped>
.card {
    background-color: #fff;
    border: 4px solid #4c324d;
    box-shadow: 8px 8px #4c324d;
    border-radius: 4px;
    margin: 20px;
    width: 640px;
    height: 480px;
}

.card-content {
    padding: 20px;
}
</style>
