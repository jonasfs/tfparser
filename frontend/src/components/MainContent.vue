<script>
import { DateTime } from 'luxon'

export default {
	name: 'MainContent',
	data() {
		return {
			headers: [
				{ text: 'Date', value: 'timestamp' },
				{ text: 'Map', value: 'mapName' },
				{ text: 'Score', value: 'score', sortable: false },
			],
			matches: {},
			sortBy: 'timestamp',
			tableKey: 0,
			fetching: 1,
			parsing: 0,
		}
	},
	methods: {
		formatDate(timestamp) {
			return DateTime.fromISO(timestamp).toRelativeCalendar()
		},
		formatScore(score1, score2) {
			const formatedScore1 = String(score1).padStart(2, "0")
			const formatedScore2 = String(score2).padStart(2, "0")
			return formatedScore1 + " - " + formatedScore2
		},
		getMatches(steamid) {
			console.log(steamid)
			global.backend.fetchMatches(steamid).then((message) => {
				console.log(message)
				this.matches = Object.values(message.payload)
			})
		},
		getPlayerProfile() {
			global.backend.getPlayerProfile().then((message) => {
				this.getMatches(message.payload.steamid)
			})
		},
	},
	created() {
		let self = this
		this.getPlayerProfile()
		global.backend.ee.on('backend-msg', (msg) => {
			if (msg.name === 'parsing') {
				self.parsing = parseInt(msg.payload)
			} else if (msg.name === 'parsed') {
				console.log("match parsed")
				//self.addMatch(msg.payload)
				self.tableKey += 1
				if (self.parsing > 0) {
					self.parsing -= 1
				}
			} else {
				console.log("debug")
				console.log(msg)
			}
		})
	},
};
</script>

<template>
	<v-main>
		<v-container
			data-test="container"
			:data-test-length="Object.values(matches).length"
			:data-test-parsing="parsing"
		>
			<v-data-table
				class="my-15"
				:headers="headers"
				:items="Object.values(matches)"
				:sort-by.sync="sortBy"
				:sort-desc="true"
				:footer-props="{
					disableItemsPerPage: true,
					itemsPerPage: 10,
					itemsPerPageText: '',
				}"
				:key="tableKey"
			>
				<template v-slot:item="{ item }">
					<tr
						:key="item.fileHash"
					>
						<td width="150"> {{ formatDate(item.timestamp) }} </td>
						<td> {{ item.mapName }} </td>
						<td> {{ formatScore(item.score1, item.score2) }} </td>
					</tr>
				</template>
			</v-data-table>
			<v-tooltip right>
				<template v-slot:activator="{ on, attrs }">
						<v-btn
							v-bind="attrs"
							v-on="on"
							class="my-10"
							elevation="3"
							v-show="parsing > 0"
							color="accent"
							small
							absolute
							fab
							top
							left
						>
							<v-badge
								:content="parsing"
								overlap
								offset-x="22"
								offset-y="22"
								dark
								color="accent"
							>
								<v-icon size="36">mdi-spin mdi-loading</v-icon>
							</v-badge>
						</v-btn>
				</template>
				<span>
					parsing {{ parsing }} matche{{ parsing != 1 ? 's' : '' }}
				</span>
			</v-tooltip>
		</v-container>
	</v-main>
</template>
