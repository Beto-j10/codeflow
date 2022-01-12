<script>

export default {
    props: {
        nodeTitle: String,
        isEmpty: Boolean,
    },
}

</script>


<template>
    <div class="node">
        <div class="node__title">{{ nodeTitle }}</div>
        <div :class="{ node__container: !isEmpty }">
            <slot></slot>
        </div>
    </div>
</template>

<style scoped>
.node__title {
    color: var(--text-color-primary);
    padding: 10px 12px;
    background-color: var(--bg-color-node-title);
    border-radius: var(--dfNodeBorderRadius) var(--dfNodeBorderRadius) 0 0;
}
.node__container {
    padding: 8px 12px 12px 12px;
}
</style>