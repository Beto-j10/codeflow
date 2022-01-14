<script>

export default {
    props: {
        nodeTitle: String,
        isEmpty: Boolean,
        nodeId: String,
    },
}

</script>


<template>
    <div class="nodeLayout">
        <div class="nodeLayout__title">{{ nodeTitle }}</div>
        <div :class="{ nodeLayout__container: !isEmpty }">
            <slot></slot>
        </div>
    </div>
</template>

<style scoped>
.nodeLayout__title {
    color: var(--text-color-primary);
    padding: 10px 12px;
    margin-bottom: 8px;
    background-color: var(--bg-color-node-title);
    border-radius: var(--dfNodeBorderRadius) var(--dfNodeBorderRadius) 0 0;
    order:  1;
    width: 100%;
    pointer-events: none;
}
.nodeLayout__container {
    margin: 8px 12px;
}
</style>