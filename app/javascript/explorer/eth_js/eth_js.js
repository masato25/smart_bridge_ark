import Vue from 'vue'
import App from './app.vue'
require("../../base/vuebase.js")

new Vue({
  el: '#app',
  components: { App }
})
console.log("vue load")
