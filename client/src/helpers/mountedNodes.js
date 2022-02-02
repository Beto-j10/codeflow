const mountedNodes = [];

export function checkMounted(nodeId) {
    return mountedNodes.includes(nodeId);
}

export function registerMounted(nodeId) {
    mountedNodes.push(nodeId);
}

export function unregisterMounted(nodeId) {
    const index = mountedNodes.indexOf(nodeId);
    if (index > -1) {
        mountedNodes.splice(index, 1);
    }
}