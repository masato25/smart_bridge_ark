import Vue from 'vue'
import ElementUI from 'element-ui'
import ifetch from '../base/ifetch'
import _ from "lodash"
require("../base/vuebase.js")

const arkExploer = "https://dexplorer.ark.io/tx"
new Vue({
  delimiters: ['@{', '}'],
  el: '#app',
  data: {
    tableData: [],
    loading: true,
  },
  mounted() {
    this.getData()
  },
  methods: {
    syncBlock() {
      this.loading = true
      ifetch("/action/ark/sync", "GET").then((res) => {
        this.getData()
      })
    },
    openExplorer(id){
      var win = window.open(`${arkExploer}/${id}`, '_blank');
      win.focus();
    },
    getData() {
      ifetch("/api/v1/data/arks.json", "GET").then((response) => {
        let data = response.data
        this.tableData = data
        this.loading = false
      })
    }
  }
})
