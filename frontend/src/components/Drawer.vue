<script>
export default {
	name: 'Drawer',
};
</script>

<template>
	<v-card>
		<v-navigation-drawer
			app
			permanent
			right
			expand-on-hover
			width="200"
			data-test="Drawer"
		>
			<v-list>
				<v-list-item class="px-2" two-line>
					<v-list-item-avatar>
						<v-img src="https://randomuser.me/api/portraits/women/81.jpg"></v-img>
					</v-list-item-avatar>

					<v-list-item-content>
						<v-list-item-title>Player</v-list-item-title>
					</v-list-item-content>
				</v-list-item>
			</v-list>

			<v-divider></v-divider>

			<v-list nav dense>
				<v-list-item link>
					<v-list-item-icon>
					</v-list-item-icon>
					<v-list-item-title>
						Trocar Usuário
					</v-list-item-title>
				</v-list-item>
			</v-list>

		</v-navigation-drawer>
	</v-card>
</template>
