<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import NodeHeader from './NodeHeader.vue';
import Node from '../layouts/Node.vue';

export default defineComponent({
    components: {
    NodeHeader,
    Node
},
})
</script>
//TODO: change Component
<template>
    <Node node-title="BlockStatement" is-empty></Node>
</template>