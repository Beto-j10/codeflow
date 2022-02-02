<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import store from '../../store';
import { division } from '../../modules/ops';
import { registerStop } from '../../helpers/stopWatch';
import Node from '../layouts/Node.vue';
import moveTitle from '../../helpers/moveTitle';
import Input from '../Input.vue';
import { checkMounted, registerMounted } from '../../helpers/mountedNodes';

export default defineComponent({
    components: {
    Node,
    Input
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
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;

            moveTitle(nodeId.value)
            if (!checkMounted(nodeId.value)) {
                registerMounted(nodeId.value)
                registerStop(nodeId.value, stop)
            }
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
        <Node node-title="Division" width="150px">
            <Input v-model.number="num" readonly />
        </Node>
    </div>
</template>