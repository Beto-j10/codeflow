<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import NodeHeader from './NodeHeader.vue';

export default defineComponent({
    components: {
        NodeHeader
    },
})
</script>
//TODO: change Component
<template>
    <div>
        <NodeHeader title="BlockStatement" />
    </div>
</template>