<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive, onUnmounted, onActivated, onRenderTracked, onRenderTriggered } from 'vue';
import NodeHeader from './NodeHeader.vue';
import store from '../../store';
import { division } from '../../modules/ops';
import { registerStop } from '../../helpers/stopWatch';
import Node from '../layouts/Node.vue';

export default defineComponent({
    components: {
        NodeHeader,
        Node
    },
    setup() {
        const el = ref(null);
        let df = null
        const nodeId = ref(0);
        const nodeData = ref({});
        const num = ref(0)
        const sharedState = reactive(store.state)
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        // console.log("parent-node: ",getCurrentInstance().appContext.app._container.__vue_app__._container.children[0].children[1].children[1].children[0].children[0].children)

        // check if the value of one of its inputs changed
        const stop = watch(sharedState, () => {
            division(df, nodeId.value)
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;
        })

        onMounted(async () => {
            await nextTick()
            console.log("Division mounted")
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;

            registerStop(nodeId.value, stop)
        });


        return {
            el,
            num,
        }
    },
})
</script>

<template>
    <div ref="el">
        <Node node-title="Division">
            <el-input-number v-model="num" :controls="false" df-num />
        </Node>
    </div>
</template>