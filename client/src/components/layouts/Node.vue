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
        <div :class="{node__container: !isEmpty}">
            <slot></slot>
        </div>
    </div>
</template>

<style scoped>
    .node__container {
        padding: 8px 12px 12px 12px;
    }
</style>