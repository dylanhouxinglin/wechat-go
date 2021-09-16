module.exports = {
    publicPath: './',
	devServer: {
		proxy: {
			'/wechat': {
				target: 'http://127.0.0.1:8989',
				changeOrigin: true,
				ws: true,
				pathRewrite: { 
					'^/wechat': '/wechat',
				}
			}
		}
	}
}