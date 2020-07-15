<template>
    <div id="nav" class="cont">
        <router-link class="home-logo" :to="this.$store.state.logged ? '/dashboard' : '/'" >
            <img src="https://svgshare.com/i/J_x.svg" alt="iban.im">
        </router-link>
    </div>
</template>

<script>
    export default {
        name: "Nav"
    }
</script>

<style scoped>

</style>