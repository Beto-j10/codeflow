<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive, onUnmounted } from 'vue';
import store from '../../store';
// import { registerStop } from '../../helpers/stopWatch';
import { checkMounted, registerMounted } from '../../helpers/mountedNodes';
import moveTitle from '../../helpers/moveTitle';
import Node from '../layouts/Node.vue';
import Input from '../Input.vue';
// import { forStructure } from '../../modules/controlStructure';

export default defineComponent({
    components: {
        Node,
        Input
    },
    setup() {
        let df = null
        const el = ref(null);
        const nodeId = ref(0);
        const nodeData = ref({});
        const num = ref(0)
        const varName = ref("")
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        onMounted(async () => {

            //TODO: finish implementation
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;
            varName.value = `For${nodeId.value}`

            nodeData.value.data.var = varName.value;
            df.updateNodeDataFromId(nodeId.value, nodeData.value.data);

            moveTitle(nodeId.value)
            if (!checkMounted(nodeId.value)) {
                registerMounted(nodeId.value)
                store.addVar(nodeId.value, varName.value)
                store.addModule(varName.value, nodeId.value, df)
            }
        });

        return {
            el,
            num,
            varName,
        }
    },
})
</script>
//TODO: change df to nodes
<template>
    <div ref="el">
        <Node :node-title="varName" width="120px" isEmpty>
            <!-- <Input v-model="varName" type="text" readonly /> -->
        </Node>
    </div>
</template>