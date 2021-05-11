const path = require('path')
// const MODE = "production";
const MODE = "development";
const enabledSourceMap = MODE === "development";
// plugin
const HtmlWebpackPlugin = require('html-webpack-plugin');
const Dotenv = require('dotenv-webpack');

module.exports = {
  entry: "./src/main.tsx",
  output: {
    path: `${__dirname}/dist`,
    filename: "main.js",
    clean: true
  },
  devServer: {
    contentBase: path.resolve(__dirname, 'dist'),
    port: 5000,
    open: true,
    historyApiFallback: true, // without no routing
  },
  module: {
    rules: [
      {
        test: /\.scss/,
        use: [
          "style-loader",
          {
            loader: "css-loader?modules", options: { url: true, sourceMap: enabledSourceMap, importLoaders: 2, },
          },
          {
            loader: "sass-loader", options: { sourceMap: enabledSourceMap, },
          },
        ],
      },
      { test: /\.(gif|png|jpg|eot|wof|woff|ttf|svg)$/, type: "asset/inline", },
      { test: /\.tsx?$/, use: "ts-loader" }
    ]
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".json"]
  },
  target: ["web", "es5"],
  plugins: [
    new HtmlWebpackPlugin({
      publicPath: '',
      filename: 'index.html',
      template: 'src/html/index.html',
    }),
    new Dotenv(),
  ]
};