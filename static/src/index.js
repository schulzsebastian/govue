import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap'
import './css/style.css'
import Vue from 'vue'
import App from './components/App.vue'

const vm = new Vue({
    el: '#app',
    components: {
        App
    }
})
window.vm = vm