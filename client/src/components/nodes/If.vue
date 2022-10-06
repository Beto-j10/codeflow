<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive, onUnmounted } from 'vue';
import store from '../../store';
import { checkMounted, registerMounted } from '../../helpers/mountedNodes';
import moveTitle from '../../helpers/moveTitle';
import Node from '../layouts/Node.vue';
import Input from '../Input.vue';
import Select from '../Select.vue';

export default defineComponent({
    components: {
        Node,
        Input,
        Select
    },
    setup() {
        let df = null
        const el = ref(null);
        const nodeId = ref("");
        const moduleName = ref("")
        const hasElse = ref(false);
        const isMounted = ref(false);

        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        onMounted(async () => {


            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            moduleName.value = `If${nodeId.value}`

            moveTitle(nodeId.value)
            hasElse.value = store.checkModuleElse(`Else${nodeId.value}`)
            if (!checkMounted(nodeId.value)) {
                registerMounted(nodeId.value)
                store.addModule(moduleName.value, nodeId.value, df)
            }
            isMounted.value = true
        });

        function handleElse(e) {
            const isElse = e.target.checked
            if (isElse) {
                store.addModuleElse(`Else${nodeId.value}`, nodeId.value, df)
            } else {
                store.removeModuleELse(nodeId.value, df)
            }
        }

        return {
            el,
            moduleName,
            handleElse,
            hasElse,
            nodeId,
            isMounted
        }
    },
})
</script>
//TODO: change df to nodes
<template>
    <div ref="el">
        <Node :node-title="moduleName" width="150px">
            <Select v-if="isMounted" :nodeID="nodeId" />
            <div class="checkElse">
                <label class="checkElse__label">Else</label>
                <input
                    class="checkElse__checkbox"
                    v-model="hasElse"
                    type="checkbox"
                    @change="handleElse"
                />
            </div>
        </Node>
    </div>
</template>

<style scoped>
.checkElse {
    display: flex;
    align-items: center;
    font-size: var(--font-size-base);
    color: var(--text-color-secondary);
    padding: 4px 0 0;
}

.checkElse__label {
    padding-right: 4px;
}

.checkElse__checkbox {
    width: 12px;
    height: 12px;
}
</style>