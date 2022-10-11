<script>

import Drawflow from 'drawflow'
import 'drawflow/dist/drawflow.min.css'
import '../styles/drawflow.css'
import { onMounted, shallowRef, h, getCurrentInstance, render, readonly, ref, nextTick, watch, reactive } from 'vue'
import Number from './nodes/Number.vue'
import NodeAdd from './nodes/Add.vue'
import NodeSub from './nodes/Subtraction.vue'
import NodeMult from './nodes/Multiplication.vue'
import NodeDiv from './nodes/Division.vue'
import NodeAssign from './nodes/Assign.vue'
import NodeFor from './nodes/For.vue'
import NodeBlock from './nodes/Block.vue'
import NodeVar from './nodes/Vars.vue'
import NodeVarFor from './nodes/VarFor.vue'
import NodeIf from './nodes/If.vue'

import store from '../store'
import { CheckClassOps, checkConnected, checkConnections, checkNodeRemoved } from '../modules/checkConnections'
import { stopWatch } from '../helpers/stopWatch'
import { transformToAST } from '../modules/transformToAST'
import BaseLayout from './layouts/BaseLayout.vue'
import { unregisterMounted } from '../helpers/mountedNodes'
import { data1, data2 } from './data'
import { importProgram } from '../modules/importProgram'


export default {
    name: "drawflow",
    setup() {
        const listNodes = readonly([
            {
                name: "Number",
                color: "#49494970",
                item: "Num",
                input: 0,
                output: 1,
                data: {
                    num: 0, // node initial value
                },
                class: "NumericLiteral",
            },
            {
                name: "Add",
                color: "#49433440",
                item: "Add",
                input: 2,
                output: 1,
                data: {
                    num: 0,
                    operator: "+",
                },
                class: "BinaryExpression Addition ops",
            },
            {
                name: "Subtraction",
                color: "#49433440",
                item: "Sub",
                input: 2,
                output: 1,
                data: {
                    num: 0,
                    operator: "-",
                },
                class: "BinaryExpression Subtraction ops",
            },
            {
                name: "Multiplication",
                color: "#49433440",
                item: "Mult",
                input: 2,
                output: 1,
                data: {
                    num: 0,
                    operator: "*",
                },
                class: "BinaryExpression Multiplication ops",
            },
            {
                name: "Division",
                color: "#49433440",
                item: "Div",
                input: 2,
                output: 1,
                data: {
                    num: 0,
                    operator: "/",
                },
                class: "BinaryExpression Division ops",
            },
            {
                name: "Assign",
                color: "#49433440",
                item: "Assign",
                input: 1,
                output: 0,
                data: {
                    num: 0,
                    var: "",
                },
                class: "VariableDeclarator Assign ops",
            },
            {
                name: "For",
                color: "#49433440",
                item: "For",
                input: 3,
                output: 0,
                data: {
                    operator: "<",
                },
                class: "ForStatement",
            },
            {
                name: "If",
                color: "#49433440",
                item: "If",
                input: 2,
                output: 0,
                data: {
                    operator: "==",
                },
                class: "IfStatement",
            },
        ]);
        const varNode = readonly(
            {
                name: "Var",
                color: "#49433440",
                item: "Var",
                input: 0,
                output: 1,
                data: {
                    num: 0,
                    idParent: 0,
                },
                class: "Identifier ops",
            },
        )
        let isModulesBar = ref(true)
        const modulesState = reactive(store.stateModules.modules)
        const varsState = reactive(store.stateVars)
        const isVarsState = ref(false)
        const editor = shallowRef({});
        const dialogVisible = ref(false);
        const dialogData = ref({});
        const Vue = { version: 3, h, render };
        const internalInstance = getCurrentInstance();
        internalInstance.appContext.app._context.config.globalProperties.$df = editor;

        function exportEditor() {
            dialogData.value = editor.value.export();
            transformToAST(dialogData.value.drawflow);
            dialogVisible.value = true;
        }
        const drag = (ev, nodeVar) => {
            ev.dataTransfer.setData("node", ev.target.getAttribute("data-node"));
            ev.dataTransfer.setData("nodeVar", nodeVar);
        };
        const drop = (ev) => {
            let data = ev.dataTransfer.getData("node");
            let nodeVar = ev.dataTransfer.getData("nodeVar");
            addNodeToDrawFlow(data, ev.clientX, ev.clientY, nodeVar);
        };
        const allowDrop = (ev) => {
        };
        function addNodeToDrawFlow(name, pos_x, pos_y, nodeVar) {
            pos_x = pos_x * (editor.value.precanvas.clientWidth / (editor.value.precanvas.clientWidth * editor.value.zoom)) - (editor.value.precanvas.getBoundingClientRect().x * (editor.value.precanvas.clientWidth / (editor.value.precanvas.clientWidth * editor.value.zoom)));
            pos_y = pos_y * (editor.value.precanvas.clientHeight / (editor.value.precanvas.clientHeight * editor.value.zoom)) - (editor.value.precanvas.getBoundingClientRect().y * (editor.value.precanvas.clientHeight / (editor.value.precanvas.clientHeight * editor.value.zoom)));
            let nodeSelected = {};
            if (nodeVar === "var") {
                nodeSelected = varNode;
                console.log("||||||||||||||||||||||||||", nodeSelected)
            } else {
                nodeSelected = listNodes.find(ele => ele.item == name);
                console.log("¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬", nodeSelected)

            }
            editor.value.addNode(name, nodeSelected.input, nodeSelected.output, pos_x, pos_y, nodeSelected.class, nodeSelected.data, name, "vue");
        };
        function handleClickItem(event, module) {
            editor.value.changeModule(module);
            const e = document.querySelectorAll(".modules__item");
            e.forEach(element => {
                element.classList.remove("modules__item--selected");
            });
            event.target.classList.add("modules__item--selected");
        }
        function toggleModulesBar() {
            isModulesBar.value = !isModulesBar.value;
        }

        function importEditor() {
            importProgram(editor.value);
        }

        function printLogs() {
            store.printStates("vars")
            store.printStates("modules")
            store.printStates("connections")
            console.log("isVarsState: ", isVarsState.value)
        }

        onMounted(async () => {
            const id = document.getElementById("drawflow");
            editor.value = new Drawflow(id, Vue, internalInstance.appContext.app._context);
            editor.value.start();
            editor.value.registerNode("Num", Number, {}, {});
            editor.value.registerNode("Add", NodeAdd, {}, {});
            editor.value.registerNode("Sub", NodeSub, {}, {});
            editor.value.registerNode("Mult", NodeMult, {}, {});
            editor.value.registerNode("Div", NodeDiv, {}, {});
            editor.value.registerNode("Assign", NodeAssign, {}, {});
            editor.value.registerNode("For", NodeFor, {}, {});
            editor.value.registerNode("Num", Number, {}, {});
            editor.value.registerNode("Block", NodeBlock, {}, {});
            editor.value.registerNode("If", NodeIf, {}, {});
            editor.value.registerNode("Var", NodeVar, {}, {});


            editor.value.on("nodeRemoved", (id) => {
                stopWatch(id);
                unregisterMounted(id);
                store.deleteState(id, editor.value)
                // checkNodeRemoved(id, editor.value);
            });
            editor.value.on("connectionCreated", (ids) => {
                // console.log("connectionCreated", ids);
                if (CheckClassOps(ids.input_id, editor.value)) {
                    store.updateConnectionsForConnectionCreated(ids, editor.value)
                }
                // checkConnected(ids.input_id, editor.value)
                // checkConnections(ids, editor.value);
                // store.updateState();
            });
            editor.value.on("connectionRemoved", (ids) => {
                // console.log("connectionRemoved", ids);
                if (CheckClassOps(ids.input_id, editor.value)) {
                    store.updateConnectionsForConnectionRemoved(ids.input_id, ids.input_class)
                }
            });

        });

        //TODO: change initial to full word
        //TODO: delete deprecated feature
        watch(varsState, () => {
            isVarsState.value = Object.keys(varsState).length > 0 ? true : false
            //     varsState.forEach(element => {
            //         if (element.name.startsWith("F")) {
            //             editor.value.registerNode(element.name, NodeVarFor, {}, {});

            //         } else if (element.name.startsWith("V")) {
            //             editor.value.registerNode(element.name, NodeVar, {}, {});
            //         }
            //     });
        });
        return {
            exportEditor,
            listNodes,
            varNode,
            drag,
            drop,
            allowDrop,
            dialogVisible,
            dialogData,
            varsState,
            isVarsState,
            modulesState,
            handleClickItem,
            isModulesBar,
            toggleModulesBar,
            importEditor,
            printLogs,
        };
    },
    components: { BaseLayout }
}

</script>

<template>
    <BaseLayout>
        <template #header>
            <h3 class="title">CodeFlow</h3>
            <el-button type="primary" @click="exportEditor">Export</el-button>
            <el-button type="primary" @click="importEditor">Import</el-button>
            <el-button type="primary" @click="printLogs">Print States</el-button>
        </template>

        <template #sidebar>
            <ul class="list">
                <!-- <details class="list__item list__item--pointer">
                    <summary>Variables</summary>
                    <ul class="vars">
                        <li class="vars__item" v-for="v in varsState" :key="v.data.idParent" draggable="true"
                            :data-node="v.item" @dragstart="drag($event, 'var')">
                            <div class="list__node">{{ v.name }}</div>
                        </li>
                    </ul>
                </details> -->
                <li class="list__item" :class="[isVarsState ? '' :'list__item--disabled']" :draggable=isVarsState
                    :data-node="varNode.item" @dragstart="drag($event, 'var')">
                    <div class="list__node">{{ varNode.name }}</div>
                </li>
                <li class=" list__item" v-for="n in listNodes" :key="n" draggable="true" :data-node="n.item"
                    @dragstart="drag($event)">
                    <!-- <div class="list__node" :style="`background: ${n.color}`">{{ n.name }}</div> -->
                    <div class="list__node">{{ n.name }}</div>
                </li>
            </ul>
        </template>

        <template #main>
            <div class="modules__container" v-if="modulesState.at(0)">
                <button class="modules__button" :class="{ 'modules__button--top': !isModulesBar }"
                    @click="toggleModulesBar">
                    <i class="modules__icon" :class="{ 'modules__icon--top': !isModulesBar }"></i>
                </button>
                <div class="modules" :class="{ 'modules--hidden': !isModulesBar }">
                    <div class="modules__bar">
                        <ul class="modules__list">
                            <li class="modules__item modules__item--selected" @click="handleClickItem($event, 'Home')">
                                Main</li>
                            <li class="modules__item" v-for="(m, i) in modulesState" :key="i"
                                @click="handleClickItem($event, m.name)">
                                <!-- <div class="modules__node" @click="changeModule(m)">{{ m }}</div> -->
                                {{ m.name }}
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
            <div id="drawflow" @drop.prevent="drop($event)" @dragover.prevent="allowDrop($event)"></div>
        </template>
    </BaseLayout>

    <el-dialog v-model="dialogVisible" title="Export" width="50%">
        <span>Data:</span>
        <pre><code>{{ dialogData }}</code></pre>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogVisible = false">Cancel</el-button>
                <el-button type="primary" @click="dialogVisible = false">Confirm</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<style scoped>
.title {
    font-size: 2.4rem;
}

.list {
    list-style: none;
    padding: 0;
    margin: 0;
    border-right: 1px solid var(--border-color-dark);
    caret-color: transparent;
}

.list__item {
    border-bottom: 1px solid var(--border-color-dark);
    cursor: move;
    padding: 12px 12px 12px var(--container-padding);
}

.list__item--pointer {
    cursor: pointer;
}

.list__item--disabled {
    color: var(--text-color-disabled);
    cursor: not-allowed;
}

.vars {
    list-style: outside;
}

.vars__item {
    cursor: move;
    margin: 8px;
}

.modules__button {
    position: absolute;
    display: block;
    top: 28px;
    left: 50%;
    height: 16px;
    width: 48px;
    margin-left: -24px;
    /* border: 1px solid var(--border-color-dark); */
    border: none;
    /* border-top: none; */
    transition-duration: 0.4s;
    background-color: transparent;
    border-radius: 0 0 8px 8px;
    cursor: pointer;
    z-index: 2;
}

.modules__button--top {
    top: 0;
    transform: rotate(180deg);
    -webkit-transform: rotate(180deg);
}

.modules__icon--top {
    transform: rotate(180deg);
    -webkit-transform: rotate(180deg);
}

.modules__icon {
    margin-top: -20px;
    border: solid var(--text-color-secondary);
    border-width: 0 2px 2px 0;
    display: inline-block;
    padding: 4px;
    transition-duration: 0.4s;
    transform: rotate(-135deg);
    -webkit-transform: rotate(-135deg);
}

/* .button:hover + .modules {
    top: 0px;
} */
/* TODO: add vars */
.modules {
    position: absolute;
    font-size: var(--font-size-small);
    color: var(--text-color-secondary);
    border-radius: 0 0 16px 16px;
    width: 80%;
    height: 44px;
    top: 0px;
    left: 10%;
    overflow: hidden;
    transition-duration: 0.4s;
    z-index: 1;
    border: 1px solid var(--border-color);
    border-top: none;
}

.modules--hidden {
    top: -44px;
}

.modules__bar {
    background-color: var(--bg-color);
    padding: 0 12px;
    overflow-x: scroll;
    overflow-y: hidden;
    scrollbar-width: thin;
    scrollbar-color: var(--scrollbar-color-thumb) var(--bg-color);
    height: 100%;
}

::-webkit-scrollbar {
    height: 2px;
}

/* Track */
::-webkit-scrollbar-track {
    background: #f1f1f1;
}

/* Handle */
::-webkit-scrollbar-thumb {
    background: var(--scrollbar-color-thumb);
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
    background: #999;
}

.modules__list {
    list-style: none;
    display: flex;
    gap: 16px;
    padding: 0;
    height: 100%;
    align-items: center;
}

.modules__item {
    cursor: pointer;
    padding: 4px 8px;
    border: 1px solid var(--border-color-light);
    border-radius: 8px;
    background-color: var(--color-white);
}

.modules__item:hover,
.modules__item--home:hover {
    background-color: var(--color-primary-light);
    color: var(--color-white);
    border-color: var(--color-primary-light);
}

.modules__item--selected {
    background-color: var(--color-primary-dark);
    color: var(--color-white);
    border-color: var(--color-primary-dark);
}

#drawflow {
    width: 100%;
    height: 100%;
    text-align: initial;
}
</style>

<style>
#drawflow .ForStatement .input_1::after {
    content: "From";
}

#drawflow .ForStatement .input_2::after {
    content: "To";
}

#drawflow .ForStatement .input_3::after {
    content: "Step";
}

#drawflow .ForStatement .output_1::before {
    content: "Do";
}

#drawflow .IfStatement .input_1::after {
    content: "Left";
}

#drawflow .IfStatement .input_2::after {
    content: "Right";
}

#drawflow .BinaryExpression .input_1::after,
#drawflow .BinaryExpression .input_2::after,
#drawflow .Assign .input_1::after,
#drawflow .Assign .output_1::before {
    content: "Number";
}

#drawflow .BinaryExpression .output_1::before {
    content: "Result";
}

/* #drawflow .Assign .input_1::after{
        content: "Variable";
    } */

#drawflow .Addition .input_1::before {
    content: "+";
}

#drawflow .Subtraction .input_1::before {
    content: "-";
}

#drawflow .Multiplication .input_1::before {
    content: "*";
}

#drawflow .Division .input_1::before {
    content: "÷";
}
</style>