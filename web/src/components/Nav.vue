<template>
    <div id="nav" class="cont">
        <v-btn rounded class="btn-home" to="/">
            Anasayfa
        </v-btn>
        <v-btn rounded v-if="true" class="btn-login" to="/login">
            Login
        </v-btn>
        <v-btn rounded v-if="false" class="btn-logout" to="/logout">
            Logout
        </v-btn>
    </div>
</template>

<script>
    export default {
        name: "Nav"
    }
</script>

<style scoped>

</style>