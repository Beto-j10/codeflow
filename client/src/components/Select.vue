<script setup>
import { ref } from 'vue';


let value = ref("--")

</script>

<template>
    <select class="select" v-model="value">
        <option value="--" disabled>Select Operator</option>
        <option value="==">Equal to (==)</option>
        <option value="!=">Not equal to (!=)</option>
        <option value=">">Greater than (&gt;)</option>
        <option value="<">Less than (&lt;)</option>
        <option value=">=">Greater than or equal to (&gt;=)</option>
        <option value="<=">Less than or equal to (&lt;=)</option>
    </select>
    <div class="select__text">
        Left <span class="select__operator"> {{ value }} </span> Right
    </div>
</template>

<style scoped>
</style>

<style scoped>
.select {
    width: 100%;
    background-color: var(--dfBackgroundColor);
    color: var(--text-color-secondary);
    border: 1px solid var(--dfNodeBorderColor);
    border-radius: var(--dfNodeBorderRadius);
    block-size: 2em;
    outline-color: var(--dfNodeBorderColor);
}
.select__text {
    font-size: var(--font-size-x-small);
    color: var(--text-color-tertiary);
    padding: 4px 0;
}
.select__operator {
    font-size: var(--font-size-small);
    /* font-weight: bold; */
}
</style>