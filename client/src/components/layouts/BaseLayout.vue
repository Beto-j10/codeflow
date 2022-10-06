<template>
    <div class="codeflow">
        <header class="header">
            <slot name="header"></slot>
        </header>
        <aside class="sidebar">
            <slot name="sidebar"></slot>
        </aside>
        <main class="main">
            <slot name="main"></slot>
        </main>
    </div>
</template>

<style scoped>
.codeflow {
    display: grid;
    grid-template-columns: 180px 1fr;
    grid-template-rows: 48px 4fr;
    grid-template-areas:
        "header header"
        "sidebar main";
    height: 100%;
}

.header {
    grid-area: header;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 var(--container-padding);
    border-bottom: 1px solid var(--border-color-dark);
    background-color: var(--bg-color);
    z-index: 2;
}

.sidebar {
    grid-area: sidebar;
}
.main {
    grid-area: main;
    padding: 0;
    position: relative;
}
</style>