<script>
import Drawer from '@/components/Drawer'
import MainContent from '@/components/MainContent'

export default {
	name: 'Main',
	components: {
		Drawer,
		MainContent,
	},
	methods: {
		refreshMain() {
			this.$refs.main.getPlayerProfile()
		}
	},
};
</script>

<template>
	<div>
		<Drawer @refreshMain="refreshMain"/>
		<MainContent
			data-test="MainContent"
			ref="main"
		/>
	</div>
</template>
