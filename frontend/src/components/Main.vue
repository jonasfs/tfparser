<script>
import Drawer from '@/components/Drawer'
import MainContent from '@/components/MainContent'

export default {
	name: 'Main',
	components: {
		Drawer,
		MainContent,
	},
};
</script>

<template>
	<div>
		<Drawer/>
		<MainContent
			data-test="MainContent"
		/>
	</div>
</template>
