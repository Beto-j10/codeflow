<script>
import NodeHeader from "./nodes/NodeHeader.vue";

export default {
    components: {
        NodeHeader
    },
    props: {
        nodeTitle: String,
    },
}

</script>


<template>
    <div class="node">
        <div class="node__title">{{ nodeTitle }}</div>
        {{refEl}}
        <div class="node__container">
            <slot></slot>
        </div>
    </div>
</template>

<style scoped>
    .node__container {
        padding: 8px 12px 12px 12px;
    }
</style>