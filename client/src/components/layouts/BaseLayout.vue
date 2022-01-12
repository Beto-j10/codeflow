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