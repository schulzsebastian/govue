var webpack = require('webpack')
module.exports = {
  entry: {
    'bundle': './src/index.js'
  },
  output: {
    filename: './dist/[name].js'
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {loader: 'babel-loader', options: {presets: ['es2015']}}
      },
      {
        test: /\.css$/,
        use: [{loader: 'style-loader'}, {loader: 'css-loader'}]
      },
      {
        test: /\.(png|jpg|gif|svg)$/,
        loader: 'file-loader',
        options: {name: './dist/fonts/[name].[ext]?[hash]'}
      },
      {
        test: /\.(otf|eot|woff|woff2|ttf|svg)$/,
        loader: 'file-loader?name=./dist/fonts/[name].[ext]'
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader'
      }
    ]
  },
  resolve: {
    alias: {vue: 'vue/dist/vue.js'}
  },
  plugins: [
    new webpack.ProvidePlugin({
      $: 'jquery',
      jQuery: 'jquery',
      'window.jQuery': 'jquery',
      Popper: ['popper.js', 'default']
    })
  ],
  devServer: {
    proxy: {
      '/': 'http://localhost:5000/'
    }
  }
}