<template>
  <div class="row justify-content-md-center">
    <div class="col-10">
      <span style="display: inline-flex;">
        <h3>Reward List</h3>
      </span>
    </div>
    <div class="col-10">
      Earn totall reward: {{reward_sum}}
      <el-table
        :data="tableData"
        style="width: 100%"
        v-loading="loading"
        :default-sort = "{prop: 'UpdatedAt', order: 'descending'}">
        <el-table-column
          prop="CreatedAt"
          label="Time"
          sortable>
          <template slot-scope="scope">
            {{converTs(scope.row.CreatedAt)}}
          </template>
        </el-table-column>
        <el-table-column
          prop="Reward"
          label="Reward">
        </el-table-column>
        <el-table-column
          prop="VoteID"
          label="Address">
        </el-table-column>
        <el-table-column
          prop="BlockID"
          label="BlockID">
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import ElementUI from 'element-ui'
import ifetch from '../../base/ifetch'
import _ from "lodash"
import moment from "moment"

const arkExploer = "https://dexplorer.ark.io/tx"

export default {
  data() {
    return {
      tableData: [],
      address: "",
      loading: true,
      reward_sum: 0,
    }
  },
  created(){
    this.getAddress()
  },
  mounted() {
    this.getData()
  },
  methods: {
    getAddress() {
      let upath = window.location.pathname
      let reg = /\/votes\/(.+)(\?.+)?/
      let address = reg.exec(upath)[1]
      this.address = address
    },
    openExplorer(id){
      var win = window.open(`${arkExploer}/${id}`, '_blank');
      win.focus();
    },
    getData() {
      ifetch(`/api/v1/data/profit/${this.address}`, "GET").then((response) => {
        let data = response.data
        this.tableData = data.reward_per_blocks
        this.reward_sum = data.reward_sum
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
