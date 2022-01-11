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
    background: #f6f6f6;
    margin-top: -15px;
    margin-left: -15px;
    margin-right: -15px;
    padding: 10px 15px;
    margin-bottom: 15px;
}
</style>