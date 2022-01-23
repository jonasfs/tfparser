<script>
import * as matchUtils from '@/plugins/matchUtils'

export default {
	name: 'MainContent',
	props: {
		userSteamid: {
			type: Number,
		},
	},
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
		addMatch(match) {
			match.score = matchUtils.getMatchScoreString(this.userSteamid, match)
			this.matches[match.fileHash] = match
		},
		getDateString(timestamp) {
			const myDate = new Date(parseInt(timestamp)*1000)
			const locale = navigator.language
			const currentYear = new Date().getFullYear()
			let dateString = ''
			dateString += new Intl.DateTimeFormat(locale, { weekday: 'short' }).format(myDate)
			dateString += ' '
			dateString += new Intl.DateTimeFormat(locale, { day: 'numeric' }).format(myDate)
			dateString += ' '
			dateString += new Intl.DateTimeFormat(locale, { month: 'short' }).format(myDate)
			if (myDate.getFullYear() < currentYear) {
				dateString += ' '
				dateString += new Intl.DateTimeFormat(locale, { year: '2-digit' }).format(myDate)
			}
			return dateString
		},
	},
	created() {
		let self = this
		global.backend.fetchMatches(this.userSteamid).then((message) => {
			Object.values(message.payload).forEach((match) => {
				self.addMatch(match)
			})
			self.tableKey += 1
			self.fetching = 0
		})
		global.backend.ee.on('backend-msg', (msg) => {
			if (msg.name === 'parsing') {
				self.parsing = parseInt(msg.payload)
			} else if (msg.name === 'parsed') {
				self.addMatch(msg.payload)
				self.tableKey += 1
				if (self.parsing > 0) {
					self.parsing -= 1
				}
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
						<td width="150"> {{ getDateString(item.timestamp) }} </td>
						<td> {{ item.mapName }} </td>
						<td> {{ item.score }} </td>
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
