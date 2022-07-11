<script>

export default {
	name: 'GreetingSetPath',
	data() {
		return {
			demoPath:
				"C:\\Program Files (x86)\\Steam\\steamapps\\common"
				+ "\\Counter-Strike Global Offensive\\csgo\\replays",
			loading: false,
			error: false,
			errorMsg: "",
		}
	},
	methods: {
		validatePath() {
			this.loading = true
			var self = this
			global.backend.validatePath(this.demoPath).then((message) => {
				self.loading = false
				console.log("validatePath return message: ")
				console.log(message)
				if (message?.name === "error") {
					self.error = true
					if (message?.payload === "os.IsNotExist") {
						self.errorMsg = "This is not a valid directory"
					} else if (message?.payload === "no files") {
						self.errorMsg = "There are no valid .dem files in this directory"
					} else if (message?.payload === "no settings") {
						self.errorMsg = "Couldn't set the demo file path on the database"
					}
				} else if (message?.name === "validatePath.callback") {
					self.error = false
					self.errorMsg = ""
					self.$emit('setParsing', message.payload)
					self.$emit('next', 2)
				}
			})
		},
	},
};
</script>

<template>
	<div>
		<p class="my-4">
			Write in the directory where your demos are stored then click GO to proceed:
		</p>
		<p class="my-4">
		<v-row>
			<v-col>
			</v-col>
			<v-col cols="9">
				<v-text-field
					dense
					prepend-icon="mdi-folder"
					v-model="demoPath"
					:disabled="loading"
					:error="error"
					:error-messages="errorMsg"
				>
				</v-text-field>
			</v-col>
			<v-col>
			</v-col>
		</v-row>
		</p>
		<p class="my-4">
			<v-btn
				color="success"
				@click="validatePath"
				:loading="loading"
				:disabled="loading"
				data-test="button"
			>
				GO
			</v-btn>
		</p>
	</div>
</template>
