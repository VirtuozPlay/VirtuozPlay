<template>
    <div>
        <button type="button" class="cursor-pointer button" @click="handleClick">
            <slot></slot>
        </button>
    </div>
</template>

<script>
    export default {
        name: 'Button',
        methods: {
            handleClick() {
                console.log("hello world")
        },
    },
};
</script>


<style lang="css" scoped>
.button {
  color: #4C324D;
  background-color: #00FF75;
  padding: 12px 24px;
  font-size: 25px;
  border-radius: 4px;
  font-weight: 700;
  border: 4px solid #4C324D;
  box-shadow: 8px 8px #4C324D;
  border-radius: 4px;
  transition: background-color 0.3s ease;
}

.button:hover {
  background-color: #FAFF00;
  color: #000;
}
</style>