const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');
const MODE = "development";

module.exports = merge(common, {
  mode: MODE,
  devtool: 'inline-source-map',
  cache: {
    type: 'filesystem',
    buildDependencies: {
      config: [__filename]
    }
  },
});