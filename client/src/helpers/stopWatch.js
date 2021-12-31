const stops = [];

/**
 * Registers the Watch Stop of the node
 * @param {string} nodeId - id of the node
 * @param {function} stop - function to stop node Watch
 */
export function registerStop(nodeId, stop) {
    stops.push({ nodeId, stop });
}

/**
 * Stop Watch manually when node is removed.
 * Watch is never stopped automatically because the component is not unmount.
 * @param {string} id - id of the node
 */
export function stopWatch(id) {
    const stp = stops.find(element => element.nodeId === id)
    if (stp) {
        stp.stop();
        stops.splice(stops.indexOf(stp), 1);
    }
}