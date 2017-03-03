const path = require("path")
const webpack = require("webpack")

module.exports = {
    entry: path.resolve("./src/main.js"),
    output: {
        path: path.resolve("./public"),
        filename: "bundle.js",
        publicPath: "/"
    },
    devtool: "cheap-module-eval-source-map",
    watch: true,
    watchOptions: {
        ignored: /node_modules/,
        poll: 1000
    },
    module: {
        rules: [
           { test: /\.js$/, loader:"babel-loader", exclude: /node_modules/ },
           { test: /\.css$/, loader: "style-loader!css-loader", exclude: /node_modules/ }
        ],
    }
}