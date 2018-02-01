import Vue from 'vue'
import ElementUI from 'element-ui'
require('jQuery')

require("../base/vuebase.js")

new Vue({
  delimiters: ['@{', '}'],
  el: '#app',
  data: {
    ark: 0,
    ethereum: 0,
    bitcoin: 0,
  },
  mounted() {
    this.ark = document.getElementById("ark").getAttribute("value")
    this.ethereum = document.getElementById("ethereum").getAttribute("value")
    this.bitcoin = document.getElementById("bitcoin").getAttribute("value")
  }
})
