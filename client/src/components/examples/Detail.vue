<template>
  <v-container fluid>
    <v-layout row>
      <v-flex xs12 sm6 offset-sm3>
        <v-card>
          <v-progress-linear :indeterminate="true" v-show="loading"></v-progress-linear>
          <v-form v-model="valid" ref="form" lazy-validation>
            <v-card-title headline>
              <h2 v-show="!editable">{{ example.title }}</h2>
              <v-text-field v-show="editable"
                label="Title"
                v-model="example.title"
                :rules="titleRules"
                :counter="60"
                required
              ></v-text-field>
            </v-card-title>
            <v-card-text>
              <p v-show="!editable">{{ example.body }}</p>
              <v-text-field v-show="editable"
                label="Body"
                v-model="example.body"
                multi-line
                :rules="bodyRules"
                required
              ></v-text-field>
            </v-card-text>
            <v-card-actions v-show="!editable">
              <v-btn small color="accent" @click="doEdit(example)">EDIT</v-btn>
              <v-spacer></v-spacer>
              <v-btn small color="error" @click="doDelete">DELETE</v-btn>
            </v-card-actions>
            <v-card-actions v-show="editable">
              <v-btn small color="success" @click="saveUpdate">SAVE</v-btn>
              <v-spacer></v-spacer>
              <v-btn small color="warning" @click="cancelEdit">CANCEL</v-btn>
            </v-card-actions>
            <v-progress-linear :indeterminate="true" v-show="processing"></v-progress-linear>
          </v-form>
        </v-card>
        <v-alert color="error" v-show="errors.length > 0" icon="warning" value="true">
          <p v-for="(error, i) in errors" :key="i">
            {{ error.message }}
          </p>
        </v-alert>
      </v-flex>
    </v-layout>
    <v-dialog v-model="showDelete">
      <v-card>
        <v-card-title class="headline">Really delete example?</v-card-title>
        <v-card-text>Are you sure? Think of the children!</v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" flat @click="cancelDelete">CANCEL</v-btn>
          <v-btn color="error" flat @click="sendDelete">DELETE</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import { HTTP } from '../../api'

export default {
  data: () => ({
    example: {},
    titleRules: [
      v => !!v || 'Title is required',
      v => (v && v.length <= 60) || 'Title mus tbe less than 60 characters'
    ],
    bodyRules: [v => !!v || 'Body is required'],
    errors: [],
    editable: false,
    showDelete: false,
    _originalExample: {},
    loading: false,
    processing: false
  }),
  methods: {
    doEdit: function(example) {
      this._originalExample = Object.assign({}, example)
      this.editable = true
    },
    cancelEdit: function() {
      Object.assign(this.example, this._originalExample)
      this.editable = false
    },
    saveUpdate: function() {
      if (this.$refs.form.validate()) {
        this.processing = true
        HTTP.put('examples/' + this.$route.params.id, this.example)
          .then(response => {
            this.editable = false
            this.processing = false
          })
          .catch(e => {
            this.errors.push(e)
            this.processing = false
          })
      }
    },
    doDelete: function() {
      this.showDelete = true
    },
    cancelDelete: function() {
      this.showDelete = false
    },
    sendDelete: function() {
      this.processing = true
      HTTP.delete('examples/' + this.$route.params.id)
        .then(response => {
          this.processing = false
          this.$router.push({ name: 'Example List' })
        })
        .catch(e => {
          this.errors.push(e)
          this.processing = false
        })
    }
  },
  created() {
    this.loading = true
    HTTP.get('examples/' + this.$route.params.id)
      .then(response => {
        this.example = response.data
        this.loading = false
      })
      .catch(e => {
        this.errors.push(e)
        this.loading = false
      })
  }
}
</script>
