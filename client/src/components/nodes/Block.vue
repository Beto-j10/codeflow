<script>
import { defineComponent, ref, onMounted, nextTick } from 'vue';
import Node from '../layouts/Node.vue';
import moveTitle from '../../helpers/moveTitle';

export default defineComponent({
    components: {
        Node
    },
    setup() {
        const el = ref(null);

        onMounted(async () => {
            await nextTick()
            const nodeId = el.value.parentElement.parentElement.id.slice(5)
            moveTitle(nodeId)
        });
        return {
            el,
        }
    }
})
</script>
//TODO: change Component
<template>
    <div ref="el">
        <Node node-title="BlockStatement" is-empty></Node>
    </div>
</template>