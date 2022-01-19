
/**
 * Move title to node root in DOM
 * @param {number} node 
 */
export default function moveTitle(node) {
    const nodeId = `node-${node}`;
    const parent = document.getElementById(nodeId)
    const element = parent.getElementsByClassName("nodeLayout__title")[0];
    const ref = parent.getElementsByClassName("inputs")[0];
    parent.insertBefore(element, ref);
}