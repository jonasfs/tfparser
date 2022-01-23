<script>
import FirstTime from '@/components/FirstTime.vue'
import Main from '@/components/Main'

export default {
	data () {
		return {
			firstTime: true,
		}
	},
	components: {
		FirstTime,
		Main
	},
	methods: {
		checkFirstTime() {
			var self = this
			global.backend.checkFirstTime().then((message) => {
				if (message.name === 'error') {
					console.log('error')
				} else {
					self.firstTime = message.payload
				}
			})
		},
		toggleFirstTime() {
			this.firstTime = !this.firstTime
		},
	},
	created() {
		this.checkFirstTime()
	}
}
</script>

<template>
  <v-container>
		<FirstTime v-if="firstTime" @toggleFirstTime="toggleFirstTime" data-test="FirstTime" />
		<Main data-test="Main" v-else />
  </v-container>
</template>

<style scoped>
</style>
