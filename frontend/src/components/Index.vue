<script>
import SplashScreen from '@/components/SplashScreen'
import PreMain from '@/components/PreMain'

export default {
	name: 'Index',
	components: {
		SplashScreen,
		PreMain,
	},
	data() {
		return {
			splash: true,
		}
	},
	created() {
		setTimeout(()=> {
			this.splash = false
		}, 2000)
	},
};
</script>

<template>
	<SplashScreen v-if="splash"/>
	<PreMain v-else />
</template>
