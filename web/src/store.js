import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import router from './router'

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        token: null,
        error: null,
    },
    mutations: {
        SET_USER_DATA(state, token) {
            localStorage.setItem('user', token);
            axios.defaults.headers.common['Authorization'] = `Bearer ${
                token
            }`;
            state.token = token
        },
        SET_ERROR(state,error){
            state.error = error;
        },
        LOGOUT() {
            localStorage.removeItem('user');
            location.reload()
        },
    },
    actions: {
        register({commit}, credentials) {
            axios.post('/graph', {
                query: `
                    mutation ($email: String!, $password: String!, $firstName: String!, $lastName: String!, $handle: String!) {
                      signUp(email: $email,password: $password, firstName: $firstName, lastName: $lastName, handle: $handle){
                        ok,
                        error,
                        user{
                          createdAt,
                          updatedAt,
                          handle,
                        },
                      }
                    }`,
                variables: {
                    ...credentials
                },
            }).then(({data}) => {
                console.log(data);
                console.log(data.data.signUp.ok);
                if(data.data.signUp.ok){
                    router.push('/login');
                }else{
                    commit('SET_ERROR', data.data.signUp.error);
                }
            }).catch( response => {
                console.log(response);
            });
        },
        login({commit}, credentials) {
            axios.post('/graph', {
                query: `
                    mutation ($email: String!, $password: String!) {
                      signIn(email: $email,password: $password){
                        ok,
                        error,
                        token,
                      }
                    }`,
                variables: {
                    ...credentials
                },
            }).then(({data}) => {
                if(data.data.signIn.ok){
                    commit('SET_USER_DATA', data.data.signIn.token);
                    router.push('/dashboard');
                }else{
                    commit('SET_ERROR', data.data.signIn.error);
                }
            }).catch( response => {
                console.log(response);
            });
        },
        logout({commit}) {
            commit('LOGOUT')
        },
    }
})