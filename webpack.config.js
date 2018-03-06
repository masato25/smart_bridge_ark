'use strict';

var ExtractTextPlugin = require("extract-text-webpack-plugin")
var path = require('path');
const goStaticPath = './app/javascript/'

module.exports = {
  cache: true,
  entry: {
    app: goStaticPath + "base/app.js",
    eth_js: goStaticPath + "explorer/eth_js/eth_js.js",
    ark_js: goStaticPath + "explorer/ark_js/ark_js.js",
    vote_js: goStaticPath + "delgates/vote_js/vote_js.js",
    vote_profit_js: goStaticPath + "delgates//vote_profit_js/vote_profit_js.js",
    pricerate_indx: goStaticPath + "pricerate/indx.js",
  },

  output: {
    path: path.resolve(__dirname, "app/assets"),
    filename: "javascript/[name].js"
  },

  module: {
    loaders: [
      {
        test: /\.css$/,
        loaders: ['style-loader','css-loader']
      },
      {
        test: /\.(scss|sass)$/,
        loader: ExtractTextPlugin.extract(["css-loader", "sass-loader"])
      },
      {
        test: /\.(png|jpg|jpeg|gif|eot|ttf|woff|woff2|svg|svgz)(\?.+)?$/,
        use: [{
          loader: 'url-loader',
          options: {
            limit: 10000
          }
        }]
      },
      {
        test: /\.(woff|woff2)(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'url-loader?limit=10000&mimetype=application/font-woff'
      },
      {
         test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
         loader: 'url-loader?limit=10000&mimetype=application/octet-stream'
      },
      {
         test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
         loader: 'file-loader'
      },
      {
        test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'url-loader?limit=10000&mimetype=image/svg+xml'
      },
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        loader: 'babel-loader',
        query: {
            presets: ['es2015']
        }
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader'
      },
      {
        test: /\.js$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
        query: {
          presets: ["es2015"]
        }
      },
    ]
  },
  resolve: {
    alias: {
      Chartist: "chartist/dist/chartist.js",
      jQuery: "jquery/dist/jquery.js",
      $: "jquery/dist/jquery.js",
      "window.jQuery":"jquery",
      Tether: 'tether',
      vue: 'vue/dist/vue.js'
    },
  },
}
