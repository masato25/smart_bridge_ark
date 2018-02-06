<template>
  <el-table
    :data="tableData"
    style="width: 100%"
    v-loading="loading"
    :default-sort = "{prop: 'UpdatedAt', order: 'descending'}">
    <el-table-column
      prop="UpdatedAt"
      label="date"
      sortable>
      <template slot-scope="scope">
        {{converTs(scope.row.UpdatedAt)}}
      </template>
    </el-table-column>
    <el-table-column
      prop="FromAddr"
      label="Sender">
    </el-table-column>
    <el-table-column
      prop="ToAddr"
      label="Receiver">
    </el-table-column>
    <el-table-column
    prop="Amount"
    label="Amount(ARK)"
    width="100">
  </el-table-column>
    <el-table-column
      prop="SmartBridge"
      label="SmartBridge">
    </el-table-column>
    <el-table-column
      prop="UpdatedAt"
      label="Time">
    </el-table-column>
    <el-table-column
      prop="ID"
      label="Explorer"
      width="100">
      <template slot-scope="scope">
        <el-button
          size="mini"
          type="primary"
          @click="openExplorer(scope.row.ID)">
          browse
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import Vue from 'vue'
import ElementUI from 'element-ui'
import ifetch from '../../base/ifetch'
import _ from "lodash"
import moment from "moment"
require("../../base/vuebase.js")

const arkExploer = "https://dexplorer.ark.io/tx"

export default {
  data() {
    return {
      tableData: [],
      loading: true,
    }
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
    },
    converTs(ts) {
      return moment(ts).format("YYYY-MM-DD HH:mm:ss")
    }
  }
}
</script>

<style>
.message {
  color: blue;
}
</style>
