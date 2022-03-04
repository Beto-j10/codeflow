<script>

export default {
    props: {
        nodeTitle: String,
        isEmpty: Boolean,
        width: String,
    },
}

</script>


<template>
    <div class="nodeLayout">
        <div class="nodeLayout__title" :style="{width: width}">{{ nodeTitle }}</div>
        <div v-if="!isEmpty" class="nodeLayout__container">
            <slot></slot>
        </div>
    </div>
</template>

<style scoped>
.nodeLayout__title {
    color: var(--text-color-secondary);
    padding: 0 12px 4px;
    background-color: var(--bg-color);
    border-radius: var(--dfNodeBorderRadius) var(--dfNodeBorderRadius) 0 0;
    order: 1;
    width: 100%;
    pointer-events: none;
    align-self: flex-start;
}

.nodeLayout__container {
    margin: 8px 12px;
}
</style>