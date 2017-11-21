<template>
  <v-container fluid>
    <v-layout row>
      <v-flex xs12 sm8 offset-sm2>
        <v-form v-model="valid" ref="form" lazy-validation>
          <v-text-field
            label="Title"
            v-model="title"
            :rules="titleRules"
            :counter="60"
            required
          ></v-text-field>
          <v-text-field
            label="Body"
            v-model="body"
            multi-line
            :rules="bodyRules"
            required
          ></v-text-field>
          <v-btn
            @click="submit"
            :disabled="!valid"
          >
            Submit
          </v-btn>
          <v-btn @click="clear">Clear</v-btn>
        </v-form>
        <v-progress-linear :indeterminate="true" v-show="sending"></v-progress-linear>
        <v-alert color="error" v-show="errors.length > 0" icon="warning" value="true">
          <p v-for="(error, i) in errors" :key="i">
            {{ error.message }}
          </p>
        </v-alert>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { HTTP } from '../../api'

export default {
  data: () => ({
    sending: false,
    title: '',
    titleRules: [
      v => !!v || 'Title is required',
      v => (v && v.length <= 60) || 'Title mus tbe less than 60 characters'
    ],
    body: '',
    bodyRules: [v => !!v || 'Body is required'],
    errors: [],
    valid: true
  }),
  methods: {
    submit() {
      if (this.$refs.form.validate()) {
        this.sending = true
        HTTP.post('examples', {
          title: this.title,
          body: this.body
        })
          .then(response => {
            this.sending = false
            this.$router.push({ name: 'Example Detail', params: { id: response.data.id } })
          })
          .catch(e => {
            this.sending = false
            this.errors.push(e)
          })
      }
    },
    clear() {
      this.$refs.form.reset()
    }
  }
}
</script>
