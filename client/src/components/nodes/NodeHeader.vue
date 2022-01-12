<template>
    <div class="header">{{ title }}</div>
</template>
<script>

export default ({
    props: {
        title: String
    },
})
</script>


<style scoped>
.header {
    background: var(--bg-color-secondary);
}
</style>