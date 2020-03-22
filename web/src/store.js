import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import router from './router'

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        token: null,
        error: null,
        isLoaded: false,
        profile: null,
    },
    mutations: {
        SET_TOKEN(state, token) {
            localStorage.setItem('user', token);
            axios.defaults.headers.common['Authorization'] = `Bearer ${
                token
            }`;
            state.token = token
        },
        SET_ERROR(state,error){
            state.error = error;
        },
        SET_USER_DATA(state, profile){
            state.profile = profile;
        },
        LOGOUT() {
            localStorage.removeItem('user');
            location.reload()
        },
        SET_IS_LOADED(state,isLoaded) {
            state.isLoaded = isLoaded
        },
        UPDATE_PROFILE(state,payload) {
            console.log(payload);
            state.profile = {
                ...state.profile,
                ...payload
            }
        }

    },
    getters: {
        getProfile: (state) => {
            return state.profile;
        }
    },
    actions: {
        fetchProfile({commit}) {
            commit('SET_IS_LOADED', false);
            axios.post('/graph',{
                query: `{
                      getMyProfile {
                        user {id,firstName,lastName,handle,bio,email},
                        ok,
                        error
                      }
                    }`,
            }).then(({data}) => {
                console.log('data');
                console.log(data);
                if(data.data.getMyProfile.ok) {
                    commit('SET_USER_DATA', data.data.getMyProfile.user);
                }else{
                    router.push('/login');
                }

            }).catch((error) => {
                console.log(error)
            }).finally(() => {
                commit('SET_IS_LOADED', true);
            })
        },
        register({commit}, credentials) {
            commit('SET_IS_LOADED', false);
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
            }).finally(() => {
                commit('SET_IS_LOADED', true);
            });
        },
        login({commit}, credentials) {
            commit('SET_IS_LOADED', false);
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
                    commit('SET_TOKEN', data.data.signIn.token);
                    router.push('/dashboard');
                }else{
                    commit('SET_ERROR', data.data.signIn.error);
                }
            }).catch( response => {
                console.log(response);
            }).finally(() => {
                commit('SET_IS_LOADED', true);
            });
        },
        logout({commit}) {
            commit('LOGOUT')
        },
    }
})