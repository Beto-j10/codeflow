<script>

export default {
    props: {
        nodeTitle: String,
        isEmpty: Boolean,
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
    background-color: var(--bg-color-node-title);
    border-radius: var(--dfNodeBorderRadius) var(--dfNodeBorderRadius) 0 0;
}
.nodeLayout__container {
    padding: 8px 12px 12px 12px;
}
</style>