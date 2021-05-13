const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');
const MODE = "production";

module.exports = merge(common, {
  mode: MODE,
});