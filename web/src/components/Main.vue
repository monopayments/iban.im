<template>
    <div>
         Shorten, create and share memorable links for IBANS.

        <b>IBAN.im/nickname/bank</b>
        <br><br>
        <router-link to="/login">
            Login
        </router-link>
        or
        <router-link to="/register" >
            Register
        </router-link>
    </div>
</template>

<script>
    export default {
        name: "Main",
        data: () => ({

        }),
        created() {
            this.$store.state.isLoaded = true
        }
    }
</script>