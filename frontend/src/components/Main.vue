<script>
import Drawer from '@/components/Drawer'
import MainContent from '@/components/MainContent'

export default {
	name: 'Main',
	components: {
		Drawer,
		MainContent,
	},
	data() {
		return {
			userSteamid: 1,
		}
	},
};
</script>

<template>
	<div>
		<Drawer/>
		<MainContent data-test="MainContent" :userSteamid="userSteamid"/>
	</div>
</template>
