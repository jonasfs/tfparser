import '@mdi/font/css/materialdesignicons.css'
import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false

if (process.env.NODE_ENV === 'production') {
	global.backend = require('./plugins/backend')

	document.addEventListener('astilectron-ready', function() {
		astilectron.onMessage((msg) => {
			global.backend.ee.emit('backend-msg', msg)
		})
		new Vue({
			vuetify,
			icons: {
				iconfont: 'mdi',
			},
			render: h => h(App)
		}).$mount('#app')
	})

	// script-src meta-tag
	var link = document.createElement('meta')
	link.setAttribute('http-equiv', 'Content-Security-Policy')
	link.setAttribute('content', 'script-src \'self\';')
	link.content = document.location
	document.getElementsByTagName('head')[0].appendChild(link);
} else if (process.env.NODE_ENV === 'development') {
	if (!window.Cypress) {
		global.backend = require('./plugins/backendMock')
	}

	new Vue({
		vuetify,
		icons: {
			iconfont: 'mdi',
		},
		render: h => h(App)
	}).$mount('#app')
}
