module.exports = {
	publicPath: process.env.NODE_ENV === 'production' ? '' : '/',
	outputDir: '../resources/app/',
	"transpileDependencies": [
		"vuetify"
	]
}
