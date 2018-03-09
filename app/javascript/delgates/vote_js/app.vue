<template>
  <div class="row justify-content-md-center">
    <div class="col-10">
      <span style="display: inline-flex;">
        <h3>Giraffe Voters</h3>
      </span>
    </div>
    <div class="col-10">
      <el-table
        :data="tableData"
        style="width: 100%"
        v-loading="loading"
        :default-sort = "{prop: 'UpdatedAt', order: 'descending'}">
        <el-table-column
          prop="CreatedAt"
          label="vote time"
          sortable>
          <template slot-scope="scope">
            {{converTs(scope.row.CreatedAt)}}
          </template>
        </el-table-column>
        <el-table-column
          prop="address"
          label="Address">
        </el-table-column>
        <el-table-column
          prop="balance"
          label="balance">
        </el-table-column>
        <el-table-column
          prop="status"
          label="status">
        </el-table-column>
        <el-table-column
          prop="weight"
          label="Weight">
        </el-table-column>
        <el-table-column
          prop="address"
          label="Profit"
          width="100">
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="primary"
              @click="openExplorer(scope.row.address)">
              browse
            </el-button>
          </template>
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
    openExplorer(id){
      var win = window.open(`/votes/${id}`, '_blank');
      win.focus();
    },
    getData() {
      ifetch("/api/v1/data/voters.json", "GET").then((response) => {
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
