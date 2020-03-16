<template>
        <v-card outlined raised class="home-content">
            <h1>Welcome to Iban.im</h1>
            <div>
                To use the app, you'll need to
                <br>
                <router-link to="/login">
                    Login
                </router-link>
                or
                <router-link to="/register" >
                    Register
                </router-link>
            </div>
        </v-card>
</template>

<script>
    export default {
        name: "Home"
    }
</script>