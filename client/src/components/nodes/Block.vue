<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import NodeHeader from './NodeHeader.vue';
import Node from '../Node.vue';

export default defineComponent({
    components: {
    NodeHeader,
    Node
},
})
</script>
//TODO: change Component
<template>
    <Node node-title="BlockStatement"></Node>
</template>